package main

import "fmt"

type maxHeap struct {
	heap []int
}

func main() {
	array := []int{3, 9, 6, 5, 4, 8, 10, 34, 2, 13, 17, 24, 1000, 56, 404}
	var maxHeap maxHeap
	fmt.Println(array)
	maxHeap.maxHeap(array)

	// fmt.Println("len", len(maxHeap.heap))
	fmt.Println(maxHeap.heap)
}

func (h *maxHeap) maxHeap(array []int) {
	h.heapify(array)
}

func (h *maxHeap) heapify(array []int) {
	h.heap = array
	for i := parent(len(h.heap) - 1); i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *maxHeap) shiftDown(currentIdx int) {
	endIdx := len(h.heap) - 1
	leftIdx := leftChild(currentIdx)

	for leftIdx <= endIdx {
		rightIdx := rightChild(currentIdx)
		idxToShift := leftIdx

		if leftIdx < endIdx && h.heap[leftIdx] < h.heap[rightIdx] {
			idxToShift = rightIdx
		} else {
			idxToShift = leftIdx
		}

		if h.heap[currentIdx] < h.heap[idxToShift] {
			h.heap[currentIdx], h.heap[idxToShift] = h.heap[idxToShift], h.heap[currentIdx]
			currentIdx = idxToShift
			leftIdx = leftChild(currentIdx)
		} else {
			return
		}
	}
}

func (h *maxHeap) shiftUp(currentIdx int) {
	parentIdx := parent(currentIdx)

	for currentIdx > 0 && h.heap[parentIdx] < h.heap[currentIdx] {
		h.heap[currentIdx], h.heap[parentIdx] = h.heap[parentIdx], h.heap[currentIdx]
		currentIdx = parentIdx
		parentIdx = parent(currentIdx)
	}
}

func (h *maxHeap) insert(value int) {
	h.heap = append(h.heap, value)
	h.shiftUp(len(h.heap) - 1)
}

func (h *maxHeap) peak() int {
	return h.heap[0]
}

func (h *maxHeap) remove() {
	h.heap[0], h.heap[len(h.heap)-1] = h.heap[len(h.heap)-1], h.heap[0]
	h.heap = h.heap[:len(h.heap)-1]
	h.shiftDown(0)
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return (i * 2) + 1
}
func rightChild(i int) int {
	return (i * 2) + 2
}
