package main

import (
	"strings"
)

type line struct {
	str          string
	indent_level uint32
}

type Compiler struct {
	code []line
}

func (self *Compiler) Compile() []string {
	lines := make([]string, len(self.code))
	for i := range lines {
		lines[i] = self.code[i].str
	}
	return lines
}

func CreateCompiler(raw string) *Compiler {
	raw_lines := strings.Split(raw, "\n")
	lines := make([]line, len(raw_lines))
	for i := range lines {
		lines[i].str = raw_lines[i]
		lines[i].indent_level = 0 // TEMPORARY
	}
	return &Compiler{lines}
}
