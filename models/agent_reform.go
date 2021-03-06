// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type agentTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *agentTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("agents").
func (v *agentTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *agentTableType) Columns() []string {
	return []string{"id", "type", "runs_on_node_id"}
}

// NewStruct makes a new struct for that view or table.
func (v *agentTableType) NewStruct() reform.Struct {
	return new(Agent)
}

// NewRecord makes a new record for that table.
func (v *agentTableType) NewRecord() reform.Record {
	return new(Agent)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *agentTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// AgentTable represents agents view or table in SQL database.
var AgentTable = &agentTableType{
	s: parse.StructInfo{Type: "Agent", SQLSchema: "", SQLName: "agents", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "Type", Type: "AgentType", Column: "type"}, {Name: "RunsOnNodeID", Type: "int32", Column: "runs_on_node_id"}}, PKFieldIndex: 0},
	z: new(Agent).Values(),
}

// String returns a string representation of this struct or record.
func (s Agent) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Type: " + reform.Inspect(s.Type, true)
	res[2] = "RunsOnNodeID: " + reform.Inspect(s.RunsOnNodeID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Agent) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Type,
		s.RunsOnNodeID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Agent) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Type,
		&s.RunsOnNodeID,
	}
}

// View returns View object for that struct.
func (s *Agent) View() reform.View {
	return AgentTable
}

// Table returns Table object for that record.
func (s *Agent) Table() reform.Table {
	return AgentTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Agent) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Agent) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Agent) HasPK() bool {
	return s.ID != AgentTable.z[AgentTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Agent) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = AgentTable
	_ reform.Struct = (*Agent)(nil)
	_ reform.Table  = AgentTable
	_ reform.Record = (*Agent)(nil)
	_ fmt.Stringer  = (*Agent)(nil)
)

type mySQLdExporterTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *mySQLdExporterTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("agents").
func (v *mySQLdExporterTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *mySQLdExporterTableType) Columns() []string {
	return []string{"id", "type", "runs_on_node_id", "service_username", "service_password", "listen_port"}
}

// NewStruct makes a new struct for that view or table.
func (v *mySQLdExporterTableType) NewStruct() reform.Struct {
	return new(MySQLdExporter)
}

// NewRecord makes a new record for that table.
func (v *mySQLdExporterTableType) NewRecord() reform.Record {
	return new(MySQLdExporter)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *mySQLdExporterTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// MySQLdExporterTable represents agents view or table in SQL database.
var MySQLdExporterTable = &mySQLdExporterTableType{
	s: parse.StructInfo{Type: "MySQLdExporter", SQLSchema: "", SQLName: "agents", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "Type", Type: "AgentType", Column: "type"}, {Name: "RunsOnNodeID", Type: "int32", Column: "runs_on_node_id"}, {Name: "ServiceUsername", Type: "*string", Column: "service_username"}, {Name: "ServicePassword", Type: "*string", Column: "service_password"}, {Name: "ListenPort", Type: "*uint16", Column: "listen_port"}}, PKFieldIndex: 0},
	z: new(MySQLdExporter).Values(),
}

// String returns a string representation of this struct or record.
func (s MySQLdExporter) String() string {
	res := make([]string, 6)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Type: " + reform.Inspect(s.Type, true)
	res[2] = "RunsOnNodeID: " + reform.Inspect(s.RunsOnNodeID, true)
	res[3] = "ServiceUsername: " + reform.Inspect(s.ServiceUsername, true)
	res[4] = "ServicePassword: " + reform.Inspect(s.ServicePassword, true)
	res[5] = "ListenPort: " + reform.Inspect(s.ListenPort, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *MySQLdExporter) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Type,
		s.RunsOnNodeID,
		s.ServiceUsername,
		s.ServicePassword,
		s.ListenPort,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *MySQLdExporter) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Type,
		&s.RunsOnNodeID,
		&s.ServiceUsername,
		&s.ServicePassword,
		&s.ListenPort,
	}
}

// View returns View object for that struct.
func (s *MySQLdExporter) View() reform.View {
	return MySQLdExporterTable
}

// Table returns Table object for that record.
func (s *MySQLdExporter) Table() reform.Table {
	return MySQLdExporterTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *MySQLdExporter) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *MySQLdExporter) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *MySQLdExporter) HasPK() bool {
	return s.ID != MySQLdExporterTable.z[MySQLdExporterTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *MySQLdExporter) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = MySQLdExporterTable
	_ reform.Struct = (*MySQLdExporter)(nil)
	_ reform.Table  = MySQLdExporterTable
	_ reform.Record = (*MySQLdExporter)(nil)
	_ fmt.Stringer  = (*MySQLdExporter)(nil)
)

type rDSExporterTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *rDSExporterTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("agents").
func (v *rDSExporterTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *rDSExporterTableType) Columns() []string {
	return []string{"id", "type", "runs_on_node_id", "listen_port"}
}

// NewStruct makes a new struct for that view or table.
func (v *rDSExporterTableType) NewStruct() reform.Struct {
	return new(RDSExporter)
}

// NewRecord makes a new record for that table.
func (v *rDSExporterTableType) NewRecord() reform.Record {
	return new(RDSExporter)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *rDSExporterTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// RDSExporterTable represents agents view or table in SQL database.
var RDSExporterTable = &rDSExporterTableType{
	s: parse.StructInfo{Type: "RDSExporter", SQLSchema: "", SQLName: "agents", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "Type", Type: "AgentType", Column: "type"}, {Name: "RunsOnNodeID", Type: "int32", Column: "runs_on_node_id"}, {Name: "ListenPort", Type: "*uint16", Column: "listen_port"}}, PKFieldIndex: 0},
	z: new(RDSExporter).Values(),
}

// String returns a string representation of this struct or record.
func (s RDSExporter) String() string {
	res := make([]string, 4)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Type: " + reform.Inspect(s.Type, true)
	res[2] = "RunsOnNodeID: " + reform.Inspect(s.RunsOnNodeID, true)
	res[3] = "ListenPort: " + reform.Inspect(s.ListenPort, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *RDSExporter) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Type,
		s.RunsOnNodeID,
		s.ListenPort,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *RDSExporter) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Type,
		&s.RunsOnNodeID,
		&s.ListenPort,
	}
}

// View returns View object for that struct.
func (s *RDSExporter) View() reform.View {
	return RDSExporterTable
}

// Table returns Table object for that record.
func (s *RDSExporter) Table() reform.Table {
	return RDSExporterTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *RDSExporter) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *RDSExporter) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *RDSExporter) HasPK() bool {
	return s.ID != RDSExporterTable.z[RDSExporterTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *RDSExporter) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = RDSExporterTable
	_ reform.Struct = (*RDSExporter)(nil)
	_ reform.Table  = RDSExporterTable
	_ reform.Record = (*RDSExporter)(nil)
	_ fmt.Stringer  = (*RDSExporter)(nil)
)

type qanAgentTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *qanAgentTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("agents").
func (v *qanAgentTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *qanAgentTableType) Columns() []string {
	return []string{"id", "type", "runs_on_node_id", "service_username", "service_password", "listen_port", "qan_db_instance_uuid"}
}

// NewStruct makes a new struct for that view or table.
func (v *qanAgentTableType) NewStruct() reform.Struct {
	return new(QanAgent)
}

// NewRecord makes a new record for that table.
func (v *qanAgentTableType) NewRecord() reform.Record {
	return new(QanAgent)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *qanAgentTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// QanAgentTable represents agents view or table in SQL database.
var QanAgentTable = &qanAgentTableType{
	s: parse.StructInfo{Type: "QanAgent", SQLSchema: "", SQLName: "agents", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "Type", Type: "AgentType", Column: "type"}, {Name: "RunsOnNodeID", Type: "int32", Column: "runs_on_node_id"}, {Name: "ServiceUsername", Type: "*string", Column: "service_username"}, {Name: "ServicePassword", Type: "*string", Column: "service_password"}, {Name: "ListenPort", Type: "*uint16", Column: "listen_port"}, {Name: "QANDBInstanceUUID", Type: "*string", Column: "qan_db_instance_uuid"}}, PKFieldIndex: 0},
	z: new(QanAgent).Values(),
}

// String returns a string representation of this struct or record.
func (s QanAgent) String() string {
	res := make([]string, 7)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Type: " + reform.Inspect(s.Type, true)
	res[2] = "RunsOnNodeID: " + reform.Inspect(s.RunsOnNodeID, true)
	res[3] = "ServiceUsername: " + reform.Inspect(s.ServiceUsername, true)
	res[4] = "ServicePassword: " + reform.Inspect(s.ServicePassword, true)
	res[5] = "ListenPort: " + reform.Inspect(s.ListenPort, true)
	res[6] = "QANDBInstanceUUID: " + reform.Inspect(s.QANDBInstanceUUID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *QanAgent) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Type,
		s.RunsOnNodeID,
		s.ServiceUsername,
		s.ServicePassword,
		s.ListenPort,
		s.QANDBInstanceUUID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *QanAgent) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Type,
		&s.RunsOnNodeID,
		&s.ServiceUsername,
		&s.ServicePassword,
		&s.ListenPort,
		&s.QANDBInstanceUUID,
	}
}

// View returns View object for that struct.
func (s *QanAgent) View() reform.View {
	return QanAgentTable
}

// Table returns Table object for that record.
func (s *QanAgent) Table() reform.Table {
	return QanAgentTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *QanAgent) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *QanAgent) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *QanAgent) HasPK() bool {
	return s.ID != QanAgentTable.z[QanAgentTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *QanAgent) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = QanAgentTable
	_ reform.Struct = (*QanAgent)(nil)
	_ reform.Table  = QanAgentTable
	_ reform.Record = (*QanAgent)(nil)
	_ fmt.Stringer  = (*QanAgent)(nil)
)

func init() {
	parse.AssertUpToDate(&AgentTable.s, new(Agent))
	parse.AssertUpToDate(&MySQLdExporterTable.s, new(MySQLdExporter))
	parse.AssertUpToDate(&RDSExporterTable.s, new(RDSExporter))
	parse.AssertUpToDate(&QanAgentTable.s, new(QanAgent))
}
