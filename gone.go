package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "1 or more arguments are required")
		return
	}

	// FIXME: Now, only first argument would be compiled.
	if !strings.HasSuffix(flag.Arg(0), ".gone") {
		fmt.Fprintln(os.Stderr, "File name extension must be \".gone\": "+flag.Arg(0))
		return
	}

	filename := flag.Arg(0)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	compiler := NewCompiler(string(content))
	err = ioutil.WriteFile(filename[:strings.LastIndex(filename, ".gone")]+".go", []byte(strings.Join(compiler.Compile(), "\n")), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
