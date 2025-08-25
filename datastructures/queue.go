// Queue implementation in Go
// FIFO data structure with enqueue, dequeue, peek, isEmpty, size, clear, and toArray methods

package datastructures

import "fmt"

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

func (q *Queue) Enqueue(element interface{}) {
	q.items = append(q.items, element)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	element := q.items[0]
	q.items = q.items[1:]
	return element
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Clear() {
	q.items = make([]interface{}, 0)
}

func (q *Queue) ToArray() []interface{} {
	result := make([]interface{}, len(q.items))
	copy(result, q.items)
	return result
}

func (q *Queue) Display() {
	fmt.Printf("Queue: %v\n", q.items)
} 