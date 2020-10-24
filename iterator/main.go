package main

import (
	"fmt"
	"go-design-patterns/iterator/bintree"
)

func main() {
	//	root := bintree.BinaryTreeNode{nil, nil, 1}
	//	root.Left = &bintree.BinaryTreeNode{nil, nil, 0}
	var tree bintree.BinaryTree
	//	tree := bintree.BinaryTree{&root}
	vals := tree.InOrderTraversalIterative()
	fmt.Println(vals)
}
