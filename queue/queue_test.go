package queue

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueue_Size(t *testing.T) {
	queue := Queue[int]{}
	for i := 0; i < 10; i++ {
		queue.Push(i)
	}
	require.Equal(t, 10, queue.Size())
}

func TestQueue_PushPopPeek(t *testing.T) {
	queue := Queue[int]{}
	value, err := queue.Pop()
	require.Error(t, err)
	require.Empty(t, value)
	value, err = queue.Peek()
	require.Error(t, err)
	require.Empty(t, value)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	value, err = queue.Peek()
	require.NoError(t, err)
	require.Equal(t, 1, value)
	value, err = queue.Pop()
	require.NoError(t, err)
	require.Equal(t, 1, value)
	value, err = queue.Peek()
	require.NoError(t, err)
	require.Equal(t, 2, value)
}
