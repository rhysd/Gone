package main

import (
	"fmt"
)

type Block struct {
	Line        uint32
	IndentLevel uint32
}

type blockNode struct {
	value Block
	next  *blockNode
}

type BlockStack struct {
	Top  *blockNode
	Size uint32
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

func (self *BlockStack) Emplace(line, indentLevel uint32) {
	self.Push(Block{line, indentLevel})
}

func (self *BlockStack) Show() {
	p := self.Top
	for p != nil {
		fmt.Println(p.value)
		p = p.next
	}
}
