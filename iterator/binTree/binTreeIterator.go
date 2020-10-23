package binTree

import (
	"go-design-patterns/iterator/iterator"
)

type binTreeIterator struct {
	stack []stackElem
}

func (iter *binTreeIterator) hasNext() bool {
	if len(iter.stack) > 0 {
		return true
	}
	return false
}

func (iter *binTreeIterator)