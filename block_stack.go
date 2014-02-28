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
	top  *blockNode
	Size int
}

func (self *BlockStack) Push(block Block) {
	self.top = &blockNode{block, self.top}
	self.Size++
}

func (self *BlockStack) Pop() {
	if self.Size > 0 {
		self.top = self.top.next
		self.Size--
	}
}

func (self *BlockStack) Top() Block {
	return self.top.value
}

func (self *BlockStack) Emplace(line, indentLevel int) {
	self.Push(Block{line, indentLevel})
}

func (self *BlockStack) IsEmpty() bool {
	return self.Size == 0
}

func (self *BlockStack) Show() {
	p := self.top
	for p != nil {
		fmt.Println(p.value)
		p = p.next
	}
}
