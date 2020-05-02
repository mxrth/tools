package integration

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

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

	//removeExecutables()
	os.Exit(r)
}

func buildExecutables() error {
	os.Mkdir("testbuilds", 0777)
	//pass ldflags to be consistent with the artifacts built by github actions
	c := exec.Command("go", "build", "-o", "testbuilds/", "-ldflags=-s -w", "./...")
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

// func RunCmd(name string, args string...) Result {
// 	exeSuffix := ""
// 	if runtime.GOOS == "windows" {
// 		exeSuffix = ".exe"
// 	}

// }

// func RunCmdWithStdin(name string, stdin io.Reader, args string...) Result {

// }

// struct Result {

// }

// func (*Result) ExitCode() {

// }

// func (*Result) WasSuccessful() {

// }

// func (*Result) Stdout() string {

// }

// func (*Result) Stderr() string {

// }
