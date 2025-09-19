package generic

import (
	"errors"
	"fmt"
)

//
// Generic Stack
//

// Stack is a generic Last-In-First-Out (LIFO) data structure
type Stack[T any] struct {
	values []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	var values []T
	return &Stack[T]{values: values}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

// Pop removes and returns the top element from the stack
// Returns an error if the stack is empty
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	l := len(s.values)
	if l <= 0 {
		return zero, errors.New("The stack is empty")
	}
	last := s.values[l-1]
	s.values = s.values[:l-1]
	return last, nil
}

// Peek returns the top element without removing it
// Returns an error if the stack is empty
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	l := len(s.values)
	if l <= 0 {
		return zero, errors.New("The stack is empty")
	}
	last := s.values[l-1]
	return last, nil
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.values)
}

// IsEmpty returns true if the stack contains no elements
func (s *Stack[T]) IsEmpty() bool {
	if len(s.values) <= 0 {
		return true
	}
	return false
}

func (s *Stack[T]) Print() {
	fmt.Println(s.values)
}

func TestStack() {
	s := &Stack[int]{}
	s.Push(1)
	s.Print()

	s.Push(2)
	s.Push(3)
	s.Print()

	s.Peek()
	s.Print()

	s.Pop()
	s.Print()
}
