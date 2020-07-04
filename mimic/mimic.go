package mimic

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"strconv"

	"xorm.io/xorm/dialects"
	"xorm.io/xorm/schemas"
)

var dsns = map[string]QueryResult{}
var counter = 0

type XormDriver struct{}

func (x *XormDriver) Parse(string, string) (*dialects.URI, error) {
	return &dialects.URI{DBType: schemas.POSTGRES}, nil
}

type QueryResult struct {
	*Result
	*Query
	NumInput int
}

type Result struct {
	NumRows int
}

type Query struct {
	Cols []string
	Vals [][]driver.Value
}

type mimic struct {
}

func (m *mimic) Open(dsn string) (driver.Conn, error) {
	if len(dsn) == 0 {
		dsn = strconv.Itoa(counter)
		counter++
	}
	return &mimicConn{dsns[dsn]}, nil
}

type mimicConn struct {
	Q QueryResult
}

func (m *mimicConn) Commit() error {
	return nil
}

func (m *mimicConn) Rollback() error {
	return nil
}

func (m *mimicConn) Prepare(query string) (driver.Stmt, error) {
	return &mimicStmt{m.Q}, nil
}

func (m *mimicConn) Close() error              { return nil }
func (m *mimicConn) Begin() (driver.Tx, error) { return m, nil }

type mimicStmt struct {
	Q QueryResult
}

func (m *mimicStmt) Close() error  { return nil }
func (m *mimicStmt) NumInput() int { return m.Q.NumInput }
func (m *mimicStmt) Exec(args []driver.Value) (driver.Result, error) {
	if m.Q.Result == nil {
		return nil, errors.New("statement was not a result type")
	}

	return &mimicResult{m.Q.Result.NumRows}, nil
}

func (m *mimicStmt) Query(args []driver.Value) (driver.Rows, error) {
	if m.Q.Query == nil {
		return nil, errors.New("statement was not a query type")
	}

	return &mimicRows{columns: m.Q.Query.Cols, values: m.Q.Query.Vals}, nil
}

type mimicResult struct {
	rowsAffected int
}

func (m *mimicResult) LastInsertId() (int64, error) {
	return 0, errors.New("not supported")
}

func (m *mimicResult) RowsAffected() (int64, error) {
	return int64(m.rowsAffected), nil
}

type mimicRows struct {
	cursor  int
	columns []string
	values  [][]driver.Value
}

func (m *mimicRows) Columns() []string { return m.columns }
func (m *mimicRows) Close() error      { return nil }
func (m *mimicRows) Next(dest []driver.Value) error {
	if m.cursor == len(m.values) {
		return io.EOF
	}

	for i, val := range m.values[m.cursor] {
		dest[i] = val
	}
	m.cursor++

	return nil
}

func init() {
	sql.Register("mimic", &mimic{})
}

func NewResult(q QueryResult) {
	dsns[strconv.Itoa(counter)] = q
}

func NewQuery(q QueryResult) {
	dsns[strconv.Itoa(counter)] = q
}

func NewResultDSN(dsn string, q QueryResult) {
	dsns[dsn] = q
}

func NewQueryDSN(dsn string, q QueryResult) {
	dsns[dsn] = q
}
