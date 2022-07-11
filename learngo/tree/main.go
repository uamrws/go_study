package main

import (
	"go_study/learngo/temp/tree"
)

func main() {
	root := new(tree.Node)
	root.Value = 1
	root.Left = &tree.Node{}
	root.Right = &tree.Node{}
	root.Left.Value = 2
	root.Right.Value = 3
	root.Left.Left = &tree.Node{4, nil, nil}
	root.Left.Right = &tree.Node{5, nil, nil}
	root.Right.Left = &tree.Node{6, nil, nil}
	root.Right.Right = &tree.Node{7, nil, nil}
	root.Traverse()
}
