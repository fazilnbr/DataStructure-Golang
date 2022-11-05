package main

import (
	"doblelinkedlist/doblelinkedlist"
	"fmt"
)

func main() {

	var n int
	fmt.Println("\nstart..................")
loop:
	for {
		fmt.Print("\n\n1--:Linked List and its operations\n\n2--:Convert Array to single linked list\n\n0--:EXIT\n\nChoose option  :-")
		fmt.Scan(&n)
		switch n {
		case 1:
			list := doblelinkedlist.Singlelinkedlist{}
			list.Start()
		case 2:
			list := doblelinkedlist.Singlelinkedlist{}
			list.Convert()
		case 0:
			break loop
		default:
			fmt.Println("Choose the correct option")

		}
	}

}
