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
	stack := BlockStack{}
	current_indent := 0
	compiled := make([]string, 0, len(self.code)*2)

	for _, l := range self.code {
		if isEmptyLine(l.str) {
			continue
		}

		if current_indent < l.indent_level {
			compiled[len(compiled)-1] += " {"
			stack.Emplace(len(compiled), l.indent_level)
		}

		compiled = append(compiled, l.str)

		if l.indent_level < current_indent {
			for !stack.IsEmpty() {
				top_block := stack.Top()
				if top_block.IndentLevel > current_indent {
					break
				}
				compiled = append(compiled, "}\n")
				stack.Pop()
			}
		}

		current_indent = l.indent_level
	}
	return compiled
}

func isEmptyLine(line string) bool {
	re := regexp.MustCompile(`^\s*(:?#.*)?$`)
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
		lines[i].indent_level = getIndentLevel(lines[i].str)
	}
	return &Compiler{lines}
}
