package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (n Node) Print() {
	fmt.Println(n.Value)
}
func (n *Node) SetValue(v int) {
	n.Value = v
}
func CreateNode(v int) *Node {
	return &Node{Value: v}
}
