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

package models

import (
	"database/sql"
	"strings"

	"github.com/go-sql-driver/mysql" // register SQL driver
	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

// Initial AUTO_INCREMENT values are spaced to prevent programming errors, or at least make them more visible.
// It does not imply that one can have at most 1000 nodes, etc.
var databaseSchema = []string{
	`CREATE TABLE schema_migrations (
		id INT NOT NULL,
		PRIMARY KEY (id)
	)`,
	`INSERT INTO schema_migrations (id) VALUES (1)`,

	`CREATE TABLE nodes (
		id INT NOT NULL AUTO_INCREMENT,
		type VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,

		region VARCHAR(255) NOT NULL DEFAULT '', -- NOT NULL for unique index below

		PRIMARY KEY (id),
		UNIQUE (type, name, region)
	) AUTO_INCREMENT = 1`,

	`INSERT INTO nodes (type, name) VALUES ('` + string(PMMServerNodeType) + `', 'PMM Server')`,

	`CREATE TABLE services (
		id INT NOT NULL AUTO_INCREMENT,
		type VARCHAR(255) NOT NULL,
		node_id INT NOT NULL,

		aws_access_key VARCHAR(255),
		aws_secret_key VARCHAR(255),
		address VARCHAR(255),
		port SMALLINT UNSIGNED,
		engine VARCHAR(255),
		engine_version VARCHAR(255),

		PRIMARY KEY (id),
		FOREIGN KEY (node_id) REFERENCES nodes (id)
	) AUTO_INCREMENT = 1000`,

	`CREATE TABLE agents (
		id INT NOT NULL AUTO_INCREMENT,
		type VARCHAR(255) NOT NULL,
		runs_on_node_id INT NOT NULL,

		service_username VARCHAR(255),
		service_password VARCHAR(255),
		listen_port SMALLINT UNSIGNED,
		qan_db_instance_uuid VARCHAR(255),

		PRIMARY KEY (id),
		FOREIGN KEY (runs_on_node_id) REFERENCES nodes (id)
	) AUTO_INCREMENT = 1000000`,

	`CREATE TABLE agent_nodes (
		agent_id INT NOT NULL,
		node_id INT NOT NULL,
		FOREIGN KEY (agent_id) REFERENCES agents (id),
		FOREIGN KEY (node_id) REFERENCES nodes (id),
		UNIQUE (agent_id, node_id)
	)`,

	`CREATE TABLE agent_services (
		agent_id INT NOT NULL,
		service_id INT NOT NULL,
		FOREIGN KEY (agent_id) REFERENCES agents (id),
		FOREIGN KEY (service_id) REFERENCES services (id),
		UNIQUE (agent_id, service_id)
	)`,
}

func OpenDB(name, username, password string, logf reform.Printf) (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = username
	cfg.Passwd = password
	cfg.DBName = name

	// todo why are below hardcoded?
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"

	// todo why we need below options?
	cfg.ClientFoundRows = true
	cfg.ParseTime = true

	dsn := cfg.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err == nil {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)
		db.SetConnMaxLifetime(0)
		err = db.Ping()
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to MySQL")
	}

	if name == "" {
		return db, nil
	}

	var count int
	if err = db.QueryRow("SELECT COUNT(*) FROM schema_migrations").Scan(&count); err == nil && count > 0 {
		return db, nil
	}

	for _, q := range databaseSchema {
		q = strings.TrimSpace(q)
		logf("\n%s\n", q)
		if _, err = db.Exec(q); err != nil {
			return nil, errors.Wrapf(err, "Failed to execute\n%s", q)
		}
	}

	return db, nil
}
