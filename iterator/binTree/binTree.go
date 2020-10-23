package binTree

type BinaryTreeNode struct {
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
	Val   int
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

type stackElem struct {
	node    BinaryTreeNode
	visited bool
}

