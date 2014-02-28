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

	content, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}

	compiler := NewCompiler(string(content))
	fmt.Println(strings.Join(compiler.Compile(), "\n"))
}
