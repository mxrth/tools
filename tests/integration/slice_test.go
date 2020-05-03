package integration

import (
	"strings"
	"testing"
)

func TestNoParams(t *testing.T) {
	r := RunCmd("slice")
	if !strings.HasPrefix(r.Stdout(), "Usage:") {
		t.Errorf("Expected usage information, got %s", r.Stdout())
	}
}
