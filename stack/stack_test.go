package stack

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack_Size(t *testing.T) {
	obj := Stack[int]{}
	for i := 0; i < 10; i++ {
		obj.Push(i)
	}
	require.Equal(t, 10, obj.Size())
}

func TestStack_Peek(t *testing.T) {
	obj := Stack[int]{}
	value, err := obj.Peek()
	require.Error(t, err)
	require.Empty(t, value)
	obj.Push(10)
	value, err = obj.Peek()
	require.NoError(t, err)
	require.Equal(t, 10, value)
}

func TestStack_Pop(t *testing.T) {
	obj := Stack[int]{
		[]int{1},
	}
	value, err := obj.Pop()
	require.NoError(t, err)
	require.Equal(t, 1, value)
	value, err = obj.Pop()
	require.Empty(t, value)
	require.Error(t, err)
}

func isComplimentaryBrackets(s1 byte, s2 byte) bool {
	return s1 == '(' && s2 == ')' ||
		s1 == '{' && s2 == '}' ||
		s1 == '[' && s2 == ']'
}

func isBalanced(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	stack := Stack[byte]{}
	for i := 0; i < len(s); i++ { // O(len(s)) = O(n)
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack.Push(s[i]) // O(1)
		} else {
			val, err := stack.Pop() // O(1)
			if err != nil || !isComplimentaryBrackets(val, s[i]) {
				return false
			}
		}
	}
	return true
}

func Test(t *testing.T) {
	inputs := []string{
		"{()}",
		"}{",
		"{([])[{[]}]}",
		"{[}]()",
		"(()",
		"))(",
	}

	for _, input := range inputs {
		if isBalanced(input) {
			fmt.Println(input, "balanced")
		} else {
			fmt.Println(input, "not balanced")
		}
	}
}
