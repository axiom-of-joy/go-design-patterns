package bintree

import (
	"errors"
)

type BinTreeIterator struct {
	stack []stackElem
}

func NewBinTreeIterator(tree BinaryTree) BinTreeIterator {
	var stack []stackElem
	if tree.Root != nil {
		stack = append(stack, stackElem{*tree.Root, false})
	}
	return BinTreeIterator{stack}
}

func (iter *BinTreeIterator) HasNext() bool {
	return len(iter.stack) > 0
}

func (iter *BinTreeIterator) GetNext() (int, error) {
	if !iter.HasNext() {
		return -1, errors.New("Cannot GetNext from empty iterator.")
	}
	foundNext := false
	var nextVal int
	for !foundNext {
		iter.stack, _, foundNext, nextVal = updateStack(iter.stack, nil)
	}
	return nextVal, nil
}
