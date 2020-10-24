package bintree

import (
	"errors"
)

type binTreeIterator struct {
	stack []stackElem
}

func NewBinTreeIterator(tree BinaryTree) binTreeIterator {
	var stack []stackElem
	if tree.Root != nil {
		stack = append(stack, stackElem{*tree.Root, false})
	}
	return binTreeIterator{stack}
}

func (iter *binTreeIterator) hasNext() bool {
	return len(iter.stack) > 0
}

func (iter *binTreeIterator) getNext() (int, error) {
	if !iter.hasNext() {
		return -1, errors.New("Tree is empty.")
	}
	foundNext := false
	var nextVal int
	for !foundNext {
		iter.stack, _, foundNext, nextVal = updateStack(iter.stack, nil)
	}
	return nextVal, nil
}
