// Min Heap implementation in Go
// Basic heap operations: insert, extractMin, peek, isEmpty, size, clear, and toArray methods

package datastructures

import (
	"fmt"
	"sort"
)

type MinHeap struct {
	heap []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		heap: make([]int, 0),
	}
}

func (h *MinHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.bubbleUp()
}

func (h *MinHeap) bubbleUp() {
	index := len(h.heap) - 1
	element := h.heap[index]

	for index > 0 {
		parentIndex := (index - 1) / 2
		parent := h.heap[parentIndex]

		if element >= parent {
			break
		}

		// Swap with parent
		h.heap[parentIndex] = element
		h.heap[index] = parent
		index = parentIndex
	}
}

func (h *MinHeap) ExtractMin() int {
	if h.IsEmpty() {
		return 0
	}

	min := h.heap[0]
	last := h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	if len(h.heap) > 0 {
		h.heap[0] = last
		h.bubbleDown()
	}

	return min
}

func (h *MinHeap) bubbleDown() {
	index := 0
	element := h.heap[0]

	for {
		leftIndex := 2*index + 1
		rightIndex := 2*index + 2
		smallest := index

		// Find smallest child
		if leftIndex < len(h.heap) && h.heap[leftIndex] < h.heap[smallest] {
			smallest = leftIndex
		}
		if rightIndex < len(h.heap) && h.heap[rightIndex] < h.heap[smallest] {
			smallest = rightIndex
		}

		if smallest == index {
			break
		}

		// Swap with smallest child
		h.heap[index] = h.heap[smallest]
		h.heap[smallest] = element
		index = smallest
	}
}

func (h *MinHeap) Peek() int {
	if h.IsEmpty() {
		return 0
	}
	return h.heap[0]
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.heap) == 0
}

func (h *MinHeap) Size() int {
	return len(h.heap)
}

func (h *MinHeap) Clear() {
	h.heap = make([]int, 0)
}

func (h *MinHeap) ToArray() []int {
	result := make([]int, len(h.heap))
	copy(result, h.heap)
	return result
}

func (h *MinHeap) Display() {
	fmt.Printf("Heap: %v\n", h.heap)
} 