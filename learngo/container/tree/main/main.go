package main

import (
	"fmt"

	"brucego.com/learngo/container/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})

	fmt.Printf("node count: %d", nodeCount)

}
