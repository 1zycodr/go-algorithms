package stack

import "fmt"

type Stack[T any] struct { // LIFO - last in, first out
	values []T
}

func (s *Stack[T]) Push(value T) { // O(1)
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (T, error) { // O(1)
	var result T
	if len(s.values) == 0 {
		return result, fmt.Errorf("stack is empty")
	}
	result = s.values[s.Size()-1]
	s.values = s.values[:s.Size()-1]
	return result, nil
}

func (s *Stack[T]) Size() int { // O(1)
	return len(s.values)
}

func (s *Stack[T]) Peek() (T, error) { // O(1)
	var result T
	if len(s.values) == 0 {
		return result, fmt.Errorf("stack is empty")
	}
	result = s.values[s.Size()-1]
	return result, nil
}
