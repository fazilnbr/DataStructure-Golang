package main

import (
	"fmt"
	"trees/trees"
)

func main() {
	var op int
	tree := trees.BST{}
	fmt.Println("\n\n		BINARY SEARCH TREE\n	      -----------------------")
loop:
	for {
		fmt.Print("\n----------------------------------------------------\n\n1--:  Insert\n\n2--:  Display\n\n3--:  Search\n\n4--:  Delete\n\n0--:  Exit\n\nPlease choose a option :-  ")
		fmt.Scan(&op)
		switch op {
		case 1:
			tree.Insert()
		case 2:
			tree.Display()
		case 3:
			fmt.Print("\nResult:  ", tree.Search())
		case 4:
			tree.Delete()
		case 0:
			break loop
		default:
			fmt.Print("\n\nChoose correct option!!!!\n\n")
		}
	}
}
