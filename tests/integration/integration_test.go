package integration

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

const BINDIR = "testbuilds"

func TestMain(m *testing.M) {
	flag.Parse()
	err := os.Chdir("../..") //we are in ./tests/integration/ now we are in .
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = buildExecutables()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := m.Run()

	removeExecutables()
	os.Exit(r)
}

func buildExecutables() error {
	os.Mkdir("testbuilds", 0777)
	//pass ldflags to be consistent with the artifacts built by github actions
	c := exec.Command("go", "build", "-o", BINDIR, "-ldflags=-s -w", "./...")
	o, err := c.CombinedOutput()
	if err != nil {
		fmt.Println(string(o))
		fmt.Println(err)
		return err
	}
	return nil
}

func removeExecutables() {
	os.RemoveAll("testbuilds/")
}

func RunCmd(name string, args ...string) Result {
	c := cmd(name, args...)
	c.Run()
	return Result{c}
}

func RunCmdWithStdin(name string, stdin io.Reader, args ...string) Result {
	c := cmd(name, args...)
	c.Stdin = stdin
	c.Run()
	return Result{c}
}

func cmd(name string, args ...string) *exec.Cmd {
	exeSuffix := ""
	if runtime.GOOS == "windows" {
		exeSuffix = ".exe"
	}

	cmd := exec.Command(fmt.Sprintf("%s%c%s%s", BINDIR, os.PathSeparator, name, exeSuffix), args...)
	cmd.Stdout = &strings.Builder{}
	cmd.Stderr = &strings.Builder{}
	return cmd
}

type Result struct {
	c *exec.Cmd
}

func (r *Result) ExitCode() int {
	return r.c.ProcessState.ExitCode()
}

func (r *Result) WasSuccessful() bool {
	return r.ExitCode() == 0
}

func (r *Result) Stdout() string {
	return r.c.Stdout.(*strings.Builder).String()
}

func (r *Result) Stderr() string {
	return r.c.Stderr.(*strings.Builder).String()
}
