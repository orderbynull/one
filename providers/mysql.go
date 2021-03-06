package providers

import (
	"database/sql"
	"fmt"
	"github.com/orderbynull/one/core"
)

const getLockQuery = "SELECT GET_LOCK('%s', 0);"
const releaseLockQuery = "DO RELEASE_LOCK('%s');"
const dsn = "%s:%s@tcp(%s:%s)/"

// MySQLLock is just struct for holding MySQL connection
type MySQLLock struct {
	db *sql.DB
}

// NewMySQLLock opens connection to MySQL.
func NewMySQLLock(user, password, host string, port string) *MySQLLock {
	db, err := sql.Open("mysql", fmt.Sprintf(dsn, user, password, host, port))
	if err != nil {
		panic(err)
	}

	return &MySQLLock{db}
}

// Lock acquires non-blocking lock via GET_LOCK().
func (m *MySQLLock) Lock(name string) bool {
	var locked bool

	err := m.db.QueryRow(fmt.Sprintf(getLockQuery, name)).Scan(&locked)
	if err != nil {
		core.ErrorAndExit(err)
	}

	return bool(locked)
}

// Unlock releasing previously acquired lock via RELEASE_LOCK()
func (m *MySQLLock) Unlock(name string) {
	if _, err := m.db.Exec(fmt.Sprintf(releaseLockQuery, name)); err != nil {
		core.ErrorAndExit(err)
	}
}

// Close closes connection to MySQL
func (m *MySQLLock) Close() {
	if err := m.db.Close(); err != nil {
		core.ErrorAndExit(err)
	}
}
