package integration

import (
	"fmt"
	"regexp"
	"strconv"
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
	exp := PassphraseRegexp(5)
	r := RunCmd("pphgen")
	if !exp.MatchString(r.Stdout()) {
		t.Errorf("Expected exactly 5 words seperated by '-' ending with a newline, got: %s", r.Stdout())
	}
	if !(r.Stderr() == "") {
		t.Errorf("Expected empty stderr, got %s", r.Stderr())
	}
}

func TestN(t *testing.T) { //TODO: test with lists
	for _, list := range []string{"", "eff", "eff_short", "de"} {
		for i := 1; i <= 10; i++ {
			exp := PassphraseRegexp(i)
			var r Result
			if list == "" {
				r = RunCmd("pphgen", "-n", strconv.FormatInt(int64(i), 10))
			} else {
				r = RunCmd("pphgen", "-n", strconv.FormatInt(int64(i), 10), "-list", list)
			}
			if !exp.MatchString(r.Stdout()) {
				t.Errorf("Expected exactly %d words seperated by '-' ending with a newline, got: %s", i, r.Stdout())
			}
			if !(r.Stderr() == "") && i > 5 { //low entropy warnings
				t.Errorf("Expected empty stderr, got %s", r.Stderr())
			}
		}
	}
}

func TestUsage(t *testing.T) {
	for _, flag := range []string{"-h", "--help", "-help"} {
		r := RunCmd("pphgen", flag)
		if !strings.HasPrefix(r.Stdout(), "Usage of") {
			t.Errorf("expected usage message when using help flag, got %s", r.Stdout())
		}
		if !(r.Stderr() == "") {
			t.Error("expected empty stderr")
		}
	}

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

func PassphraseRegexp(n int) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf("^([[:alnum:]]+-){%d}[[:alnum:]]+\n$", n-1))
}
