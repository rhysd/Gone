package main

import (
	"regexp"
	"strings"
	"unicode"
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

func isEmptyLine(line string) bool {
	re := regexp.MustCompile(`^\\s*(:?#.*)?$`)
	return re.MatchString(line)
}

func indentLevel(line string) uint32 {
	var level uint32 = 0
	for _, r := range line {
		if !unicode.IsSpace(r) {
			return level
		}
		switch r {
		case '\t':
			level += 8
		case ' ':
			level += 1
		}
	}
	return level
}

func NewCompiler(raw string) *Compiler {
	raw_lines := strings.Split(raw, "\n")
	lines := make([]line, len(raw_lines))
	for i := range lines {
		lines[i].str = raw_lines[i]
		lines[i].indent_level = indentLevel(lines[i].str)
	}
	return &Compiler{lines}
}
