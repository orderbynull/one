package core

import (
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"time"
)

const statsTmpl = "------------\nStats: exit code=%d; duration=%.1f sec.\n"

var errorCannotLock = errors.New("lock already acquired")

// execute runs command and returns some metrics.
func execute(command string) (int, float64) {
	var exitCode int

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout

	start := time.Now()

	if err := cmd.Run(); err != nil {
		exitCode = 1
	} else {
		exitCode = 0
	}

	return exitCode, time.Since(start).Seconds()
}

// Process holds all lock/exec/unlock logic for any provider.
func Process(locker Locker, name string, cmd string) error {
	defer locker.Free()

	if locker.Lock(name) {
		defer locker.Unlock(name)
		execute(cmd)

		return nil
	} else {
		return errorCannotLock
	}
}
