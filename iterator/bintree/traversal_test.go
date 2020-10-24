package bintree

import (
	"fmt"
	"testing"
)

func slicesEqual(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, elem := range slice1 {
		if elem != slice2[i] {
			return false
		}
	}
	return true
}

func TestInOrderTraversal(t *testing.T) {

	// Set up.
	tree0 := BinaryTree{nil}
	var expectedSlice0 []int = nil

	root1 := BinaryTreeNode{nil, nil, 5}
	tree1 := BinaryTree{&root1}
	expectedSlice1 := []int{5}

	root2 := BinaryTreeNode{nil, nil, 5}
	root2.Left = &BinaryTreeNode{nil, nil, 4}
	tree2 := BinaryTree{&root2}
	expectedSlice2 := []int{4, 5}

	root3 := BinaryTreeNode{nil, nil, 5}
	root3.Right = &BinaryTreeNode{nil, nil, 6}
	tree3 := BinaryTree{&root3}
	expectedSlice3 := []int{5, 6}

	root4 := BinaryTreeNode{nil, nil, 5}
	root4.Left = &BinaryTreeNode{nil, nil, 3}
	root4.Left.Right = &BinaryTreeNode{nil, nil, 4}
	tree4 := BinaryTree{&root4}
	expectedSlice4 := []int{3, 4, 5}

	root5 := BinaryTreeNode{nil, nil, 5}
	root5.Right = &BinaryTreeNode{nil, nil, 7}
	root5.Right.Left = &BinaryTreeNode{nil, nil, 6}
	tree5 := BinaryTree{&root5}
	expectedSlice5 := []int{5, 6, 7}

	var tests = []struct {
		tree          BinaryTree
		expectedSlice []int
	}{
		{tree0, expectedSlice0},
		{tree1, expectedSlice1},
		{tree2, expectedSlice2},
		{tree3, expectedSlice3},
		{tree4, expectedSlice4},
		{tree5, expectedSlice5}}

	// Execute test cases.
	var test_name string
	for i, test := range tests {
		test_name = fmt.Sprintf("TestInOrderTraversalIterative-%d", i)
		t.Run(test_name, func(t *testing.T) {
			actual := test.tree.InOrderTraversalIterative()
			if !slicesEqual(actual, test.expectedSlice) {
				t.Fatal("Expected", test.expectedSlice, "Actual", actual)
			}
		})
		test_name = fmt.Sprintf("TestInOrderTraversalRecursive-%d", i)
		t.Run(test_name, func(t *testing.T) {
			actual := test.tree.InOrderTraversalRecursive()
			if !slicesEqual(actual, test.expectedSlice) {
				t.Fatal("Expected", test.expectedSlice, "Actual", actual)
			}
		})
	}
}
