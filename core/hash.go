package core

import (
	"fmt"
	"hash/fnv"
)

// MakeLockName returns hash-based lock name for given command.
func MakeLockName(cmd string) string {
	h := fnv.New32a()
	h.Write([]byte(cmd))

	return fmt.Sprintf("%d_lock", h.Sum32())
}
