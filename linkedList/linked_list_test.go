package linkedList

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNode_InsertAfterAndSize(t *testing.T) {
	list := Node[int]{
		Value: 1,
	}
	list.InsertAfter(2)
	size := list.Size()
	require.Equal(t, 2, size)
	require.NotNil(t, list.Next)
	require.Equal(t, 2, list.Next.Value)
}

func TestNode_RemoveAfter(t *testing.T) {
	list := Node[int]{
		Value: 1,
	}
	list.InsertAfter(2)
	list.InsertAfter(3)
	list.RemoveAfter()
	require.Equal(t, 2, list.Size())
	require.Equal(t, 2, list.Next.Value)
}

func TestNode_Access(t *testing.T) {
	list := Node[int]{
		Value: 1,
	}
	list.InsertAfter(2)
	list.InsertAfter(3)

	node := list.Access(3)
	require.NotNil(t, node)
	require.Equal(t, 3, node.Value)
	require.Equal(t, 2, node.Next.Value)

	node = list.Access(4)
	require.Nil(t, node)
}
