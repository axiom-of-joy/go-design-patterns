package bintree

func (tree *BinaryTree) InOrderTraversalRecursive() []int {
	root := tree.Root
	var vals []int
	root.InOrderTraversalRecursive(&vals)
	return vals
}

func (node *BinaryTreeNode) InOrderTraversalRecursive(vals *[]int) {
	if node == nil {
		return
	} else {
		node.Left.InOrderTraversalRecursive(vals)
		*vals = append(*vals, node.Val)
		node.Right.InOrderTraversalRecursive(vals)
	}
}

func updateStack(stack []stackElem, vals []int) ([]stackElem, []int, bool, int) {
	elem := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	node := elem.node
	visited := elem.visited
	var lastVal int
	appended := false
	if visited {
		vals = append(vals, node.Val)
		lastVal = node.Val
		appended = true
		if node.Right != nil {
			elem := stackElem{*node.Right, false}
			stack = append(stack, elem)
		}
	} else {
		elem := stackElem{node, true}
		stack = append(stack, elem)
		if node.Left != nil {
			elem := stackElem{*node.Left, false}
			stack = append(stack, elem)
		}
	}
	return stack, vals, appended, lastVal
}

func (tree *BinaryTree) InOrderTraversalIterative() []int {
	var vals []int
	if tree.Root == nil {
		return vals
	}
	var rootElem = stackElem{*tree.Root, false}
	var stack = []stackElem{rootElem}
	for len(stack) > 0 {
		stack, vals, _, _ = updateStack(stack, vals)
	}
	return vals
}
