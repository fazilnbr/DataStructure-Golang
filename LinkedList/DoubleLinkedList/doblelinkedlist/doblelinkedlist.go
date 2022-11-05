package doblelinkedlist

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}
type Singlelinkedlist struct {
	head *Node
	tail *Node
}

func (list *Singlelinkedlist) Start() {
	var n int
	fmt.Print("\n\n                         SINGLE LINKED LIST\n                       ----------------------\n\n\n")
loop:
	for {
		fmt.Print("Enter the options:-\n\n1-->Insert\n2-->Display\n3-->Delete Last\n4-->Delete Specific\n5-->Insertbeg\n6-->sort\n7-->Delete Duplicates\n8-->display raverce\n0-->Exit\n\n")
		fmt.Print("\n\nEnter a option :-  ")
		fmt.Scan(&n)
		switch n {
		case 1:
			list.insert(0)
		case 2:
			list.display()
		case 3:
			list.deletelast()
		case 4:
			list.deletespecific()
		case 5:
			list.insertbeg()
		case 6:
			list.sort()
		case 7:
			list.deleteDuplicate()
		case 8:
			list.displayReverce()
		case 0:
			break loop
		default:
			fmt.Println("\nChoose the correct option................!")

		}
	}
}

func (list *Singlelinkedlist) insert(a int) {
	var data int
	if a == 0 {
		fmt.Print("\n\n\nEnter the number:-  ")
		fmt.Scan(&data)
	} else {
		data = a
	}
	newNode := new(Node)
	newNode.data = data
	newNode.prev = nil
	newNode.next = nil

	if list.head == nil {
		list.head = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
	}

	list.tail = newNode

}

func (list *Singlelinkedlist) insertbeg() {
	var data int
	fmt.Print("\n\n\nEnter the number:-  ")
	fmt.Scan(&data)
	newNode := new(Node)
	newNode.data = data
	newNode.next = list.head
	list.head.prev = newNode

	list.head = newNode

	if list.tail == nil {
		list.tail = list.head
	}
}

func (list *Singlelinkedlist) display() {
	if list.head == nil {
		fmt.Println("The list is empty...........!")
	} else {
		fmt.Print("The list is  :-  ")
		temp := list.head
		for temp != nil {
			fmt.Print("  ", temp.data)
			temp = temp.next
		}
		fmt.Println()
	}
}

func (list *Singlelinkedlist) displayReverce() {
	if list.head == nil {
		fmt.Println("The list is empty...........!")
	} else {
		fmt.Print("The list is  :-  ")
		temp := list.tail
		for temp != nil {
			fmt.Print("  ", temp.data)
			temp = temp.prev
		}
		fmt.Println()
	}
}

func (list *Singlelinkedlist) deletelast() {
	if list.head == nil {
		fmt.Println("There is no data to delete...........!")
	} else {
		if list.head == list.tail {
			list.head = nil
			list.tail = nil
		} else {
			list.tail = list.tail.prev
			list.tail.next = nil
		}
		fmt.Println("deleted")
	}
}

func (list *Singlelinkedlist) deletespecific() {
	if list.head == nil {
		fmt.Println("There is no data to delete...........!")
	} else {

		var data int
		fmt.Print("\n\n\nEnter the number:-  ")
		fmt.Scan(&data)
		temp := list.head
		if temp.data == data {
			list.head = temp.next
			list.head.prev = nil
		} else {

			for temp != nil {
				if temp.next.data == data {
					temp.next = temp.next.next
					break
				}
				temp = temp.next
			}
		}
	}

}

func (list *Singlelinkedlist) sort() {
	temp := list.head

	for temp != nil {
		temp2 := list.head
		for temp2 != nil {
			if temp.data < temp2.data {
				t := temp.data
				temp.data = temp2.data
				temp2.data = t
			}
			temp2 = temp2.next
		}
		temp = temp.next
	}

}

func (list *Singlelinkedlist) deleteDuplicate() {
	if list.head == nil {
		fmt.Println("There is no data to delete...........!")
	} else {
		temp := list.head

		for temp != nil {
			temp2 := temp
			for temp2 != nil {
				if temp.data == temp2.next.data {
					temp2.next = temp2.next.next
					break
				}
				temp2 = temp2.next
			}
			temp = temp.next
		}

	}
}

func (list *Singlelinkedlist) Convert() {
	var n int

	fmt.Print("\n\nEnter the the size of array  :-")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("\n\nEnter elaements  :-")
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Println("\nThe array you enterd  :  ", arr, "\n\nLet convert to Linked List   :-  ")
	for _, val := range arr {
		list.insert(val)
	}
	fmt.Println("After the convertion  :  ")
	list.display()
}
