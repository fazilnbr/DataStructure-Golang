package Actions

import (
	"fmt"
)

func Start() {
	var n int
	fmt.Print("\n\n                         SINGLE LINKED LIST\n                       ----------------------\n\n\n")
loop:
	for {
		fmt.Print("Enter the options:-\n\n1-->Insert\n2-->Display\n3-->Delete\n4-->Exit\n\n")
		fmt.Print("\n\nEnter a option :-  ")
		fmt.Scan(&n)
		switch n {
		case 1:
			insert()
		case 2:
			display()
		case 4:
			break loop
		default:
			fmt.Println("\nChoose the correct option................!")

		}
	}
}

func insert() {
	fmt.Println("inserted..........")
}

func display() {
	fmt.Println("display")
}
