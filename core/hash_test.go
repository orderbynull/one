package core

import (
	"github.com/nu7hatch/gouuid"
	"testing"
)

func TestMakeLockNameForUniqueness(t *testing.T) {
	generatedLocks := make(map[string]string)

	for i := 0; i < 1000000; i++ {
		if u, err := uuid.NewV4(); err == nil {
			name, _ := MakeLockName(u.String())

			if name == "" {
				t.Fatal("Got \"\" for lock name, expected not empty")
			}

			if _, exists := generatedLocks[name]; exists {
				t.Errorf("Got non-unique lock name on iteration #%d", i)
			} else {
				generatedLocks[name] = u.String()
			}
		}
	}
}

func TestMakeLockNameWithEmptyCmd(t *testing.T) {
	name, err := MakeLockName("")
	if err == nil {
		t.Error("Got nil, error expected")
	}

	if name != "" {
		t.Error("Got non empty lock name, \"\" expected")
	}
}
