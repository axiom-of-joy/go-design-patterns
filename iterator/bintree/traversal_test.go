package bintree

import (
	"fmt"
	"testing"
)

type testEntry struct {
	tree          BinaryTree
	expectedSlice []int
}

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

func getTestEntries() []testEntry {
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

	var tests = []testEntry{
		{tree0, expectedSlice0},
		{tree1, expectedSlice1},
		{tree2, expectedSlice2},
		{tree3, expectedSlice3},
		{tree4, expectedSlice4},
		{tree5, expectedSlice5}}
	return tests
}

func TestTraversals(t *testing.T) {

	// Set up.
	tests := getTestEntries()

	// Execute test cases.
	var testName string
	for i, test := range tests {
		testName = fmt.Sprintf("TestInOrderTraversalIterative-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := test.tree.InOrderTraversalIterative()
			if !slicesEqual(actual, test.expectedSlice) {
				t.Fatal("Expected", test.expectedSlice, "Actual", actual)
			}
		})
		testName = fmt.Sprintf("TestInOrderTraversalRecursive-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := test.tree.InOrderTraversalRecursive()
			if !slicesEqual(actual, test.expectedSlice) {
				t.Fatal("Expected", test.expectedSlice, "Actual", actual)
			}
		})
		testName = fmt.Sprintf("TestBinTreeIterator-%d", i)
		t.Run(testName, func(t *testing.T) {
			iter := NewBinTreeIterator(test.tree)
			var actual []int
			for iter.hasNext() {
				nextVal, _ := iter.getNext()
				actual = append(actual, nextVal)
			}
			if !slicesEqual(actual, test.expectedSlice) {
				t.Fatal("Expected", test.expectedSlice, "Actual", actual)
			}
		})
	}
}
