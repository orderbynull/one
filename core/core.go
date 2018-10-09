package core

import (
	"fmt"
	"github.com/orderbynull/one/providers"
	"os"
	"os/exec"
	"time"
)

const statsTmpl = "------------\nStats: exit code=%d; duration=%.1f sec.\n"
const lockAcquired = "Lock '%s' already acquired. Exiting."

// execute ...
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

// Process ...
func Process(locker providers.Locker, name string, cmd string) {
	defer locker.Free()

	var exitCode int
	var duration float64

	if locker.Lock(name) {
		defer locker.Unlock()
		exitCode, duration = execute(cmd)
	} else {
		println(fmt.Sprintf(lockAcquired, name))
	}

	fmt.Printf(statsTmpl, exitCode, duration)
}
