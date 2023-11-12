package queue

import "fmt"

type Queue[T any] struct {
	values []T
}

func (q *Queue[T]) Push(value T) { // O(1)
	q.values = append(q.values, value)
}

func (q *Queue[T]) Size() int { // O(1)
	return len(q.values)
}

func (q *Queue[T]) Pop() (T, error) { // O(1)
	var result T
	if q.Size() == 0 {
		return result, fmt.Errorf("queue is empty")
	}
	result = q.values[0]
	q.values = q.values[1:q.Size()]
	return result, nil
}

func (q *Queue[T]) Peek() (T, error) {
	var result T
	if q.Size() == 0 {
		return result, fmt.Errorf("queue is empty")
	}
	result = q.values[0]
	return result, nil
}
