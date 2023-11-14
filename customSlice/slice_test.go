package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInit(t *testing.T) {
	slice := Init(1, 2, 3)
	require.NotEmpty(t, slice)
}

func TestSlice_GetAndSet(t *testing.T) {
	slice := Init(1, 2, 3)
	value, err := slice.Get(0)
	require.NoError(t, err)
	require.Equal(t, 1, value)
	err = slice.Set(0, 2)
	require.NoError(t, err)
	value, err = slice.Get(0)
	require.NoError(t, err)
	require.Equal(t, 2, value)
	// check error cases
	value, err = slice.Get(10)
	require.Error(t, err)
	require.Empty(t, value)
	err = slice.Set(10, 2)
	require.Error(t, err)
}

func TestAppend(t *testing.T) {
	slice := Init(1, 2, 3)
	slice, err := Append(slice, 4)
	require.NoError(t, err)
	require.Equal(t, 4, slice.Len())
	require.Equal(t, 6, slice.Cap())
	value, err := slice.Get(3)
	require.NoError(t, err)
	require.Equal(t, 4, value)
}

func TestBigSliceAppend(t *testing.T) {
	slice := Init(1)
	var err error
	for i := 1; i < 1000; i++ {
		slice, err = Append(slice, i)
		require.NoError(t, err)
	}
	require.Equal(t, 1000, slice.Len())
	require.Equal(t, 1232, slice.Cap())
}

func TestSliceAppendManyValues(t *testing.T) {
	slice := Init(1)
	appendSlice := make([]int, 999, 999)
	for i := 0; i < 999; i++ {
		appendSlice[i] = i + 2
	}
	slice, err := Append(slice, appendSlice...)
	require.NoError(t, err)
	require.Equal(t, 1000, slice.Len())
	require.Equal(t, 1000, slice.Cap())
}
