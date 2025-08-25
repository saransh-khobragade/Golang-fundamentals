// Linked List implementation in Go
// Singly linked list with append, prepend, delete, find, reverse, and utility methods

package datastructures

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		size: 0,
	}
}

func (l *LinkedList) Append(data interface{}) {
	newNode := &Node{Data: data, Next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.size++
}

func (l *LinkedList) Prepend(data interface{}) {
	newNode := &Node{Data: data, Next: l.head}
	l.head = newNode
	l.size++
}

func (l *LinkedList) Delete(data interface{}) bool {
	if l.head == nil {
		return false
	}

	if l.head.Data == data {
		l.head = l.head.Next
		l.size--
		return true
	}

	current := l.head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			l.size--
			return true
		}
		current = current.Next
	}
	return false
}

func (l *LinkedList) Find(data interface{}) *Node {
	current := l.head
	for current != nil {
		if current.Data == data {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.head
	var next *Node

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.head = prev
}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) Peek() interface{} {
	if l.head != nil {
		return l.head.Data
	}
	return nil
}

func (l *LinkedList) GetAt(index int) interface{} {
	if index < 0 || index >= l.size {
		return nil
	}

	current := l.head
	count := 0

	for count < index {
		current = current.Next
		count++
	}

	return current.Data
}

func (l *LinkedList) InsertAt(data interface{}, index int) bool {
	if index < 0 || index > l.size {
		return false
	}

	if index == 0 {
		l.Prepend(data)
		return true
	}

	newNode := &Node{Data: data, Next: nil}
	current := l.head
	count := 0

	for count < index-1 {
		current = current.Next
		count++
	}

	newNode.Next = current.Next
	current.Next = newNode
	l.size++
	return true
}

func (l *LinkedList) RemoveAt(index int) interface{} {
	if index < 0 || index >= l.size {
		return nil
	}

	if index == 0 {
		data := l.head.Data
		l.head = l.head.Next
		l.size--
		return data
	}

	current := l.head
	count := 0

	for count < index-1 {
		current = current.Next
		count++
	}

	data := current.Next.Data
	current.Next = current.Next.Next
	l.size--
	return data
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.size = 0
}

func (l *LinkedList) ToArray() []interface{} {
	result := make([]interface{}, l.size)
	current := l.head
	index := 0

	for current != nil {
		result[index] = current.Data
		current = current.Next
		index++
	}

	return result
}

func (l *LinkedList) Display() {
	fmt.Printf("LinkedList: %v\n", l.ToArray())
} 