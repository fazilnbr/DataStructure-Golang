// this program on linked list
package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}
type Queue struct {
	head *node
	tail *node
}

func (Queue *Queue) Start() {
	var n int
	fmt.Print("\n\n                         SINGLE LINKED LIST\n                       ----------------------\n\n\n")
loop:
	for {
		fmt.Print("Enter the options:-\n\n1-->EnQueue\n2-->Display\n3-->DeQueue\n0-->Exit\n\n")
		fmt.Print("\n\nEnter a option :-  ")
		fmt.Scan(&n)
		switch n {
		case 1:
			Queue.enqueue(0)
		case 2:
			Queue.display()
		case 3:
			Queue.dequeue()
		case 0:
			break loop
		default:
			fmt.Println("\nChoose the correct option................!")

		}
	}
}

func (Queue *Queue) enqueue(a int) {
	var data int
	if a == 0 {
		fmt.Print("\n\n\nEnter the number:-  ")
		fmt.Scan(&data)
	} else {
		data = a
	}
}

func (Queue *Queue) display() {
	if Queue.head == nil {
		fmt.Println("The list is empty...........!")
	} else {
		fmt.Print("The list is  :-  ")
		temp := Queue.head
		for temp != nil {
			fmt.Println("  ", temp.data)
			temp = temp.next
		}
		fmt.Println()
	}
}
func (Queue *Queue) dequeue() {
	if Queue.head == nil {
		fmt.Println("There is no data to delete...........!")
	} else {
		temp := Queue.head
		for temp != nil {
			if temp.next.next == nil {
				temp.next = nil
				Queue.tail = temp
				break
			}
			temp = temp.next
		}
	}
}
