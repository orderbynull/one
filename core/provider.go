package core

import (
	_ "github.com/go-sql-driver/mysql"
)

// Locker is base interface for all lock providers.
// Every provider must implement it in order to be supported by application.
type Locker interface {
	Lock(name string) bool
	Unlock(name string)
}

// Closer is interface for disposable resources
type Closer interface {
	Close()
}
