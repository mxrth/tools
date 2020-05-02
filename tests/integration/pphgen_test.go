package integration

import (
	"strings"
	"testing"
)

func TestNoArgs(t *testing.T) {
	r := RunCmd("pphgen")
	if !r.WasSuccessful() {
		t.Error("Could not run pphgen")
	}
}

func TestDefaultN(t *testing.T) {

}

func TestWrongList(t *testing.T) {
	r := RunCmd("pphgen", "-list", "efff")
	if r.WasSuccessful() {
		t.Error("executing pphgen with wrong list shouldn't be successful error code")
	}
	if !strings.HasPrefix(r.Stderr(), "ERROR:") {
		t.Error("Executing pphgen with a wrong list should signal an error")
	}
	if r.Stdout() != "" {
		t.Error("Executing pphgen with wrong list should not produce any output")
	}
}
