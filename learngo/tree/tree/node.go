package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
func (node *Node) Print() {
	fmt.Println(node.Value)
}
