package main

import (
	"fmt"
)

type Block struct {
	Line        int
	IndentLevel int
}

type blockNode struct {
	value Block
	next  *blockNode
}

type BlockStack struct {
	Top  *blockNode
	Size int
}

func (self *BlockStack) Push(block Block) {
	self.Top = &blockNode{block, self.Top}
	self.Size++
}

func (self *BlockStack) Pop() {
	if self.Size > 0 {
		self.Top = self.Top.next
		self.Size--
	}
}

func (self *BlockStack) Emplace(line, indentLevel int) {
	self.Push(Block{line, indentLevel})
}

func (self *BlockStack) Show() {
	p := self.Top
	for p != nil {
		fmt.Println(p.value)
		p = p.next
	}
}
