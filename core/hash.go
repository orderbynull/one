package core

import (
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
)

var errorEmptyCmd = errors.New("empty cmd")

// MakeLockName returns base64 lock name for given command.
func MakeLockName(cmd string) (string, error) {
	if cmd != "" {
		return fmt.Sprintf("%s_lock", base64.StdEncoding.EncodeToString([]byte(cmd))), nil
	} else {
		return "", errorEmptyCmd
	}
}
