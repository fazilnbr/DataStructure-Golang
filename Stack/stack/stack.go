package stack

import (
	"fmt"
)

type node struct {
	prev *node
	data int
}

type Stack struct {
	top *node
}

func (Stack *Stack) Start() {
	var n int
	fmt.Print("\n\n                         SINGLE LINKED LIST\n                       ----------------------\n\n\n")
loop:
	for {
		fmt.Print("\nEnter the options:-\n\n1-->PUSH\n2-->POP\n3-->DISPLAY\n0-->Exit\n\n")
		fmt.Print("\n\nEnter a option :-  ")
		fmt.Scan(&n)
		switch n {
		case 1:
			Stack.push()
		case 2:
			Stack.pop()
		case 3:
			Stack.display()
		case 0:
			break loop
		default:
			fmt.Println("\nChoose the correct option................!")

		}
	}
}

func (Stack *Stack) push() {
	var data int
	fmt.Print("\n\nEnter number to push      :")
	fmt.Scan(&data)

	newNode := new(node)
	newNode.prev = nil
	newNode.data = data

	if Stack.top == nil {
		Stack.top = newNode
	} else {
		newNode.prev = Stack.top
		Stack.top = newNode
	}
}

func (Stack *Stack) pop() {
	if Stack.top == nil {
		fmt.Println("The stack is empty...................!")
	} else {
		Stack.top = Stack.top.prev
	}
}

func (Stack *Stack) display() {
	if Stack.top == nil {
		fmt.Println("The stack is empty...................!")
	} else {
		fmt.Print("\n\nThe Stack is  :-  ")
		temp := Stack.top
		for temp != nil {
			fmt.Print(" ", temp.data)
			temp = temp.prev
		}
	}
}
