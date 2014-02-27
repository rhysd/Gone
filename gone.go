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
		fmt.Fprintln(os.Stderr, "required 1 or more arguments")
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
