package integration

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	// err := buildExecutables()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	os.Exit(m.Run())
}

func buildExecutables() error {
	err := os.Chdir("../..") //we are in ./tests/integration/ now we are in .
	if err != nil {
		return err
	}

	os.Mkdir("testbuilds", os.ModeDir)
	//run go build -o testbuilds/ -race
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
