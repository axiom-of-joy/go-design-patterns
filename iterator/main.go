package main

import (
	"fmt"
	"go-design-patterns/iterator/bintree"
)

func main() {
	root := bintree.BinaryTreeNode{nil, nil, 3}
	root.Left = &bintree.BinaryTreeNode{nil, nil, 1}
	root.Left.Right = &bintree.BinaryTreeNode{nil, nil, 2}
	tree := bintree.BinaryTree{&root}
	iter := bintree.NewBinTreeIterator(tree)
	for iter.HasNext() {
		nextVal, _ := iter.GetNext()
		fmt.Printf("%d", nextVal)
	}
}
