package main

import (
	"regexp"
	"strings"
	"unicode"
)

type line struct {
	str          string
	indent_level int
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

func getIndentLevel(line string) int {
	var level int = 0
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
		lines[i].indentLevel = indentLevel(lines[i].str)
	}
	return &Compiler{lines}
}
