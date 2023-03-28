package main

import "fmt"

type NodeTree struct {
	value                       int
	leftNextNode, rightNextNode *NodeTree
}

func createNode(value int) *NodeTree {
	return &NodeTree{value: value}
}

func printFunc() func(node *NodeTree) {
	return func(node *NodeTree) {
		fmt.Print("\t", node.value)
	}
}

func sumFunc() func(node *NodeTree) {
	sum := 0
	return func(node *NodeTree) {
		sum += node.value
		fmt.Println(sum)
	}
}

// 可以使用入参为闭包函数，方便实现不同的功能
func (nodeTree *NodeTree) traverseFunc(f func(node *NodeTree)) {
	if nil == nodeTree {
		return
	}
	nodeTree.leftNextNode.traverseFunc(f)
	f(nodeTree)
	nodeTree.rightNextNode.traverseFunc(f)
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

	//					5
	//			3			 8
	//		1	  nil	 4		0
	pf := printFunc()
	sf := sumFunc()

	fmt.Println("===printFUN")
	root.traverseFunc(pf)
	fmt.Println()
	fmt.Println("===sumFun")
	root.traverseFunc(sf)
	fmt.Println("===getMaxNodeWithChannel")
	c := root.traverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.value > maxNode {
			maxNode = node.value
		}
	}
	fmt.Printf("MaxNodeWithChannel=%d\n", maxNode)
}

func (nodeTree *NodeTree) traverseWithChannel() chan *NodeTree {
	c := make(chan *NodeTree)
	go func() {
		nodeTree.traverseFunc(func(node *NodeTree) {
			c <- node
		})
		close(c)
	}()
	return c
}
