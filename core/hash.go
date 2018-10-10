package core

import (
	"fmt"
	"hash/fnv"
)

func MakeLockName(cmd string) string {
	h := fnv.New32a()
	h.Write([]byte(cmd))

	return fmt.Sprintf("%s_lock", h.Sum32())
}
