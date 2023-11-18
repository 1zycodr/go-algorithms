package main

import "fmt"

type Slice[T any] struct {
	array []T // imagine that its array
	len   int
	cap   int
}

func Init[T any](values ...T) Slice[T] {
	sliceLen := len(values)
	sliceCap := sliceLen

	return Slice[T]{
		len:   sliceLen,
		cap:   sliceCap,
		array: values,
	}
}

func (s *Slice[T]) Len() int {
	return s.len
}

func (s *Slice[T]) Cap() int {
	return s.cap
}

func (s *Slice[T]) Get(i int) (T, error) {
	var result T
	if s.len <= i {
		return result, fmt.Errorf("index out of range [%d] with length %d", i, s.len)
	}
	return s.array[i], nil
}

func (s *Slice[T]) Set(i int, value T) error {
	if s.len <= i {
		return fmt.Errorf("index out of range [%d] with length %d", i, s.len)
	}
	s.array[i] = value
	return nil
}

func Append[T any](s Slice[T], values ...T) (Slice[T], error) {
	neededCap := len(values) + s.Len()
	oldCap := s.Cap()
	newCap := oldCap
	doubleCap := oldCap + oldCap
	if neededCap > oldCap {
		// newCap calculation
		if neededCap > doubleCap {
			newCap = neededCap
		} else {
			const threshold = 256 // like in 1.20.5 go implementation
			if oldCap < threshold {
				newCap = doubleCap
			} else {
				for newCap < neededCap {
					newCap += (newCap + 3*threshold) / 4
				}
			}
		}
		// copy old array and add new values
		newArray := make([]T, neededCap, newCap)
		for i, val := range s.array {
			newArray[i] = val
		}
		for i, val := range values {
			newArray[i+oldCap] = val
		}
		s.array = newArray
		s.cap = newCap
		s.len = neededCap
	} else {
		s.len = neededCap
		// new allocations never called, imagine that it is pointer reassignment
		s.array = append(s.array, values...)
	}
	return s, nil
}

func (s *Slice[T]) Cut(i, j int) (*Slice[T], error) {
	if i >= j {
		return nil, fmt.Errorf("invalid indexes")
	}
	result := s.array[i:j]
	return &Slice[T]{
		array: result,
		len:   len(result),
		cap:   cap(result),
	}, nil
}
