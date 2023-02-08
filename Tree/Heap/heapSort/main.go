package main

import "fmt"

var array []int

func main() {
	insert()

	fmt.Println(array)

	sort()

	fmt.Println(array)

}

func insert() {
	array = append(array, 3, 9, 6, 5, 4, 8, 10, 34, 2, 13, 17, 24, 1, 56)
}

func sort() {
	sortHelper(len(array))
}

func sortHelper(length int) {
	if length-1 < 0 {
		return
	}
	heapify(length)

	array[0], array[length-1] = array[length-1], array[0]

	sortHelper(length - 1)

}

func heapify(length int) {
	for i := parent(length - 1); i >= 0; i-- {
		shiftDown(i, length)
	}
}

func shiftDown(currentIdx int, length int) {
	endIdx := length - 1
	leftIdx := leftChild(currentIdx)

	for leftIdx <= endIdx {
		rightIdx := rightChild(currentIdx)
		idxToShift := leftIdx

		if rightIdx <= endIdx && array[leftIdx] < array[rightIdx] {
			idxToShift = rightIdx
		} else {
			idxToShift = leftIdx
		}

		if array[currentIdx] < array[idxToShift] {
			array[currentIdx], array[idxToShift] = array[idxToShift], array[currentIdx]
			currentIdx = idxToShift
			leftIdx = leftChild(currentIdx)
		} else {
			return
		}
	}
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
