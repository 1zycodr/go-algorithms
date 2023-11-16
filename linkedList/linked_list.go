package linkedList

type Comparable interface {
	comparable
}

type Node[T Comparable] struct {
	Value T
	Next  *Node[T]
}

func (n *Node[T]) InsertAfter(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  n.Next,
	}
	n.Next = newNode
}

func (n *Node[T]) RemoveAfter() {
	if n.Next != nil {
		n.Next = n.Next.Next
	}
}

func (n *Node[T]) Access(value T) *Node[T] {
	head := n
	for head != nil {
		if head.Value == value {
			return head
		}
		head = head.Next
	}
	return nil
}

func (n *Node[T]) Size() int {
	result := 0
	head := n
	for head != nil {
		result++
		head = head.Next
	}
	return result
}
