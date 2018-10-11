package core

import (
	"fmt"
	"github.com/orderbynull/one/providers"
	"os"
	"os/exec"
	"time"
)

const statsTmpl = "------------\nStats: exit code=%d; duration=%.1f sec.\n"
const lockAcquired = "Lock already acquired. Exiting."

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
func Process(locker providers.Locker, name string, cmd string) {
	defer locker.Free()

	var exitCode int
	var duration float64

	if locker.Lock(name) {
		defer locker.Unlock(name)
		exitCode, duration = execute(cmd)
		fmt.Printf(statsTmpl, exitCode, duration)
	} else {
		println(lockAcquired)
	}
}
