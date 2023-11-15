package linkedList

type Node[T any] struct {
	Value T
	Next  *T
}
