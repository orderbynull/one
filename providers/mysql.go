package providers

import (
	"database/sql"
	"fmt"
	"log"
)

type MySQLLock struct {
	db   *sql.DB
	lock string
}

// NewMySQLLock ...
func NewMySQLLock() *MySQLLock {
	db, err := sql.Open("mysql",
		"root:515528aA@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err)
	}

	return &MySQLLock{db, ""}
}

// Lock ...
func (m *MySQLLock) Lock(name string) bool {
	var locked bool

	m.lock = name
	err := m.db.QueryRow(fmt.Sprintf("SELECT GET_LOCK('%s', 0);", name)).Scan(&locked)
	if err != nil {
		log.Fatal(err)
	}

	return bool(locked)
}

// Unlock ...
func (m *MySQLLock) Unlock() {
	m.db.Exec(fmt.Sprintf("DO RELEASE_LOCK('%s');", m.lock))
}

// Free ...
func (m *MySQLLock) Free() {
	m.db.Close()
}
