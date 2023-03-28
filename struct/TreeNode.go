package main

import (
	"fmt"
)

type NodeTree struct {
	value                       int
	leftNextNode, rightNextNode *NodeTree
}

func createNode(value int) *NodeTree {
	return &NodeTree{value: value}
}

func (nodeTree NodeTree) print() {
	fmt.Println(nodeTree.value)
}

func (nodeTree *NodeTree) setValue(value int) {
	if nil == nodeTree {
		fmt.Println("nil pointer！")
		return
	}
	nodeTree.value = value
}

func (nodeTree *NodeTree) traverse() {
	if nil == nodeTree {
		return
	}
	nodeTree.leftNextNode.traverse()
	fmt.Print("\t", nodeTree.value)
	nodeTree.rightNextNode.traverse()
}

func main() {
	root := NodeTree{5, nil, nil}
	root.leftNextNode = &NodeTree{3, nil, nil}
	root.rightNextNode = &NodeTree{8, nil, nil}
	root.rightNextNode.leftNextNode = &NodeTree{value: 4}
	root.rightNextNode.rightNextNode = createNode(0)
	nodeTree := new(NodeTree)
	nodeTree.value = 1
	root.leftNextNode.leftNextNode = nodeTree

	root.leftNextNode.print()
	root.leftNextNode.rightNextNode.setValue(33)
	//					5
	//			3			 8
	//		1	  nil	 4		0(33)
	//
	root.rightNextNode.rightNextNode.setValue(33)
	root.rightNextNode.rightNextNode.print()
	root.traverse()
}
