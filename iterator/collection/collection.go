package collection

import (
	"go-design-patterns/iterator/iterator"
)

type collection interface {
	createIterator() iterator.Iterator
}
