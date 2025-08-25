// Stack implementation in Go
// LIFO data structure with push, pop, peek, isEmpty, size, clear, and toArray methods

package datastructures

import "fmt"

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

func (s *Stack) Push(element interface{}) {
	s.items = append(s.items, element)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	element := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return element
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = make([]interface{}, 0)
}

func (s *Stack) ToArray() []interface{} {
	result := make([]interface{}, len(s.items))
	copy(result, s.items)
	return result
}

func (s *Stack) Display() {
	fmt.Printf("Stack: %v\n", s.items)
} 