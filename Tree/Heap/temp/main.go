package main

import (
	"fmt"
)

type MaxHeap struct {
	arr []int
}

func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return (2 * i) + 2
}
func parent(i int) int {
	return (i - 1) / 2
}

func main() {
	data := []int{34, 45, 12, 451, 67, 123, 2, 1678, 343, 22}

	heap := MaxHeap{}
	for i := 0; i < len(data); i++ {
		heap.insert(data[i])
	}
	fmt.Println(len(heap.arr))
	fmt.Println(heap.arr)
	heap.rootDelete()
	fmt.Println(heap.arr)

}
func (h *MaxHeap) insert(value int) {
	h.arr = append(h.arr, value)
	h.insertHelper(len(h.arr) - 1)

}
func (h *MaxHeap) insertHelper(position int) {
	if h.arr[position] > h.arr[(position-1)/2] {
		swap(h.arr, position, (position-1)/2)
		h.insertHelper((position - 1) / 2)

	}

}
func (h *MaxHeap) peek() {
	for i, v := range h.arr {
		fmt.Println(i, " ", v)
	}
}
func swap(arr []int, a, b int) {
	//fmt.Println("swap", a, b)
	arr[a], arr[b] = arr[b], arr[a]

}
func (h *MaxHeap) rootDelete() {

	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.rootDeleteHelper(0)

}
func (h *MaxHeap) maxValue(position int) int { //45, 22, 12, 451, 67, 123, 2, 1678, 343,
	//fmt.Println(left(position),right(position),"max value")

	if h.arr[left(position)] > h.arr[right(position)] {

		return left(position)

	} else if h.arr[left(position)] < h.arr[right(position)] {
		return right(position)
	}
	return left(position)

}

func (h *MaxHeap) rootDeleteHelper(position int) {
	if right(position) >= len(h.arr) {
		return
	}
	w := h.maxValue(position)
	fmt.Println(w)
	if h.arr[position] < h.arr[w] {
		swap(h.arr, position, w)
		fmt.Println(h.maxValue(position), position, w, "mac positinh")
		h.rootDeleteHelper(w)
	}

}
func (h *MaxHeap) buildheap(a []int) {
	h.arr = a
	fmt.Println(h.arr)
	for i := len(h.arr)/2 - 1; i >= 0; i-- {
		fmt.Println(i)
		MaxHeapify(h.arr, i)

	}

}
func MaxHeapify(heap []int, i int) {
	left := left(i)
	right := right(i)
	large := i
	n := len(heap) - 1
	if left <= n && heap[left] > heap[large] {
		large = left
	}
	if right <= n && heap[right] > heap[large] {
		large = right
	}
	fmt.Println(large)
	if large != i {
		fmt.Println(heap, large, i)
		swap(heap, large, i)
		fmt.Println(heap)
		MaxHeapify(heap, large)
	}

}
