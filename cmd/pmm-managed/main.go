// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	_ "expvar"
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/Percona-Lab/pmm-managed/api"
	"github.com/Percona-Lab/pmm-managed/handlers"
	"github.com/Percona-Lab/pmm-managed/service"
)

const (
	shutdownTimeout = 3 * time.Second
)

var (
	// TODO combine gRPC and REST ports?
	gRPCAddrF  = flag.String("listen-grpc-addr", "127.0.0.1:7771", "gRPC server listen address")
	restAddrF  = flag.String("listen-rest-addr", "127.0.0.1:7772", "REST server listen address")
	debugAddrF = flag.String("listen-debug-addr", "127.0.0.1:7773", "Debug server listen address")

	// certFileF = flag.String("cert-file", "cert.pem", "TLS certificate file for gRPC server")
	// keyFileF  = flag.String("key-file", "key.pem", "TLS key file for gRPC server")

	prometheusConfigF = flag.String("prometheus-config", "", "Prometheus configuration file path")
	prometheusURLF    = flag.String("prometheus-url", "http://127.0.0.1:9090/", "Prometheus base URL")
)

// runGRPCServer runs gRPC server until context is canceled, then gracefully stops it.
func runGRPCServer(ctx context.Context) {
	l := logrus.WithField("component", "gRPC")
	l.Infof("Starting server on http://%s/ ...", *gRPCAddrF)

	prometheusURL, err := url.Parse(*prometheusURLF)
	if err != nil {
		l.Panic(err)
	}
	gRPCServer := grpc.NewServer()
	server := &handlers.Server{
		Prometheus: &service.Prometheus{
			ConfigPath: *prometheusConfigF,
			URL:        prometheusURL,
		},
	}
	api.RegisterBaseServer(gRPCServer, server)
	api.RegisterAlertsServer(gRPCServer, server)

	listener, err := net.Listen("tcp", *gRPCAddrF)
	if err != nil {
		l.Panic(err)
	}
	go func() {
		for {
			err = gRPCServer.Serve(listener)
			if err == grpc.ErrServerStopped {
				break
			}
			l.Errorf("Failed to serve: %s", err)
		}
		l.Info("Server stopped.")
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	go func() {
		<-ctx.Done()
		gRPCServer.Stop()
	}()
	gRPCServer.GracefulStop()
	cancel()
}

// runGRPCServer runs REST proxy server until context is canceled, then gracefully stops it.
func runRESTServer(ctx context.Context) {
	l := logrus.WithField("component", "REST")
	l.Infof("Starting server on http://%s/ ...", *restAddrF)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	type registrar func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error
	for _, r := range []registrar{
		api.RegisterBaseHandlerFromEndpoint,
		api.RegisterAlertsHandlerFromEndpoint,
	} {
		if err := r(ctx, mux, *gRPCAddrF, opts); err != nil {
			l.Panic(err)
		}
	}

	server := &http.Server{
		Addr:     *restAddrF,
		ErrorLog: log.New(os.Stderr, "runRESTServer: ", 0),
		Handler:  mux,

		// TODO we probably will need it for TLS+HTTP/2, see https://github.com/philips/grpc-gateway-example/issues/11
		// TLSConfig: &tls.Config{
		// 	NextProtos: []string{"h2"},
		// },
	}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			l.Panic(err)
		}
		l.Info("Server stopped.")
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	if err := server.Shutdown(ctx); err != nil {
		l.Errorf("Failed to shutdown gracefully: %s", err)
	}
	cancel()
}

// runGRPCServer runs debug server until context is canceled, then gracefully stops it.
func runDebugServer(ctx context.Context) {
	l := logrus.WithField("component", "debug")
	msg := `Starting debug server ...
            pprof    http://%s/debug/pprof
            expvar   http://%s/debug/vars
            requests http://%s/debug/requests
            events   http://%s/debug/events`
	l.Infof(msg, *debugAddrF, *debugAddrF, *debugAddrF, *debugAddrF)

	server := &http.Server{
		Addr:     *debugAddrF,
		ErrorLog: log.New(os.Stderr, "runDebugServer: ", 0),
	}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			l.Panic(err)
		}
		l.Info("Server stopped.")
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	if err := server.Shutdown(ctx); err != nil {
		l.Errorf("Failed to shutdown gracefully: %s", err)
	}
	cancel()
}

type v2Logger struct {
	*logrus.Entry
}

func (v *v2Logger) V(l int) bool {
	return true
}

// check interface
var _ grpclog.LoggerV2 = (*v2Logger)(nil)

func main() {
	log.SetFlags(0)
	log.SetPrefix("stdlog: ")
	logrus.SetLevel(logrus.DebugLevel)
	grpclog.SetLoggerV2(&v2Logger{logrus.WithField("component", "grpclog")})
	flag.Parse()

	l := logrus.WithField("component", "main")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer l.Info("Done.")

	// handle termination signals: first one gracefully, force exit on the second one
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-signals
		l.Warnf("Got %v (%d) signal, shutting down...", s, s)
		cancel()

		s = <-signals
		l.Panicf("Got %v (%d) signal, exiting!", s, s)
	}()

	// start servers, wait for them to exit
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		runGRPCServer(ctx)
	}()
	go func() {
		defer wg.Done()
		runRESTServer(ctx)
	}()
	go func() {
		defer wg.Done()
		runDebugServer(ctx)
	}()
	wg.Wait()
}
