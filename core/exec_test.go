package core

import "testing"

func TestExecute(t *testing.T) {
	code, duration := execute("sleep 2")
	if duration < 2 {
		t.Errorf("Expected < %d, got %f", 2, duration)
	}
	if code > 0 {
		t.Errorf("Expected %d, got %d", 0, code)
	}

	code, _ = execute("non-existing")
	if code != 1 {
		t.Errorf("Expected %d, got %d", 1, code)
	}
}
