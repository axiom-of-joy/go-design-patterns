package main

import (
	"fmt"
	"go-design-patterns/iterator/binTree"
)

func main() {
	//	root := binTree.BinaryTreeNode{nil, nil, 1}
	//	root.Left = &binTree.BinaryTreeNode{nil, nil, 0}
	var tree binTree.BinaryTree
	//	tree := binTree.BinaryTree{&root}
	vals := tree.InOrderTraversalIterative()
	fmt.Println(vals)
}
