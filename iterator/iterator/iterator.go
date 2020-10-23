package iterator

type iterator interface {
	hasNext() bool
	getNext() int
}
