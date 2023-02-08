package main

import "fmt"

type minHeap struct {
	heap []int
}

func main() {
	array := []int{9, 5, 1}
	var minHeap minHeap
	minHeap.MinHeap(array)

	fmt.Println("len", len(minHeap.heap))
	fmt.Println(minHeap.heap)
	minHeap.insert(21)
	minHeap.insert(8)
	minHeap.insert(18)
	minHeap.insert(12)
	minHeap.insert(2)
	minHeap.insert(36)
	fmt.Println(minHeap.heap)
	minHeap.remove()
	fmt.Println(minHeap.heap)
	minHeap.remove()
	fmt.Println(minHeap.heap)
	fmt.Println("peak", minHeap.peak())

}

func (h *minHeap) MinHeap(array []int) {
	h.buildHeap(array)

}

func (h *minHeap) buildHeap(array []int) {
	h.heap = array
	for i := parent(len(h.heap) - 1); i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *minHeap) shiftDown(currentIdx int) {
	endIdx := len(h.heap) - 1
	leftIdx := leftChild(currentIdx)

	for leftIdx <= endIdx {
		rightIdx := rightChild(currentIdx)
		var idxToShift int

		if rightIdx <= endIdx && h.heap[rightIdx] < h.heap[leftIdx] {
			idxToShift = rightIdx
		} else {
			idxToShift = leftIdx
		}
		if h.heap[currentIdx] > h.heap[idxToShift] {
			h.heap[currentIdx], h.heap[idxToShift] = h.heap[idxToShift], h.heap[currentIdx]
			currentIdx = idxToShift
			leftIdx = leftChild(currentIdx)
		} else {
			return
		}
	}

}

func (h *minHeap) shiftUp(currentIdx int) {
	parentIdx := parent(currentIdx)

	for currentIdx > 0 && h.heap[parentIdx] > h.heap[currentIdx] {

		h.heap[currentIdx], h.heap[parentIdx] = h.heap[parentIdx], h.heap[currentIdx]

		currentIdx = parentIdx
		parentIdx = parent(currentIdx)
	}
}

func (h *minHeap) insert(value int) {
	h.heap = append(h.heap, value)
	fmt.Println("inserted", h.heap[len(h.heap)-1])
	h.shiftUp(len(h.heap) - 1)

}

func (h *minHeap) peak() int {
	return h.heap[0]
}

func (h *minHeap) remove() {
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
