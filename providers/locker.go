package providers

import (
	_ "github.com/go-sql-driver/mysql"
)

type Locker interface {
	Lock(name string) bool
	Unlock(name string)
	Free()
}
