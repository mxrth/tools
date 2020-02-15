package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]

	outf, err := os.Create(name + ".go")
	if err != nil {
		panic(err)

	}

	fileName := name + ".txt"
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")
	fmt.Fprintf(outf, "package main\n\n")
	fmt.Fprintf(outf, "var %s  = []string {\n", name)
	for _, l := range lines {
		w := strings.Split(l, "\t")
		fmt.Fprintf(outf, "\"%s\",\n", w[1])
	}
	fmt.Fprintln(outf, "}")
}
