package main

import "fmt"

const size = 3

type hashmap struct {
	arr [size]bucket
}

type bucket struct {
	head *node
	tail *node
}
type node struct {
	data string
	next *node
}

// insert funtion of hash map
func (mp *hashmap) insert(key *string) {
	newnode := new(node)
	newnode.data = *key
	newnode.next = nil
	index := hash(key)
	if mp.arr[index].head == nil {
		mp.arr[index].head = newnode
		mp.arr[index].tail = newnode
	} else {
		mp.arr[index].tail.next = newnode
		mp.arr[index].tail = newnode
	}
}

func listdisplay(head *node) {
	if head == nil {
		fmt.Println("The list is empty...........!")
	} else {
		// fmt.Print("The list is  :-  ")
		temp := head
		for temp != nil {
			fmt.Print("  ", temp.data)
			temp = temp.next
		}
		fmt.Println()
	}
}

func (mp *hashmap) display() {
	for _, v := range mp.arr {
		if v.head != nil {
			listdisplay(v.head)
		}
	}
}

func listsearch(head *node, key *string) bool {
	if head == nil {
		fmt.Println("The list is empty...........!")
	} else {
		temp := head
		for temp != nil {
			if temp.data == *key {
				return true
			}
			temp = temp.next
		}
	}
	return false
}

func (mp *hashmap) search(key *string) bool {
	index := hash(key)
	fmt.Println("serach index:", index)
	if mp.arr[index].head == nil {
		return false
	} else {
		return listsearch(mp.arr[index].head, key)
	}

}

func (mp *hashmap) listdelete(index int, key *string) {

	temp := mp.arr[index].head
	if temp.data == *key {
		mp.arr[index].head = temp.next
		fmt.Println("Data deleted")
	} else {

		for temp != nil {
			if temp.next.data == *key {
				if temp.next.next == nil {
					mp.arr[index].tail = temp
				}
				temp.next = temp.next.next
				fmt.Println("Data deleted")
				break
			}
			temp = temp.next
		}
	}

}

func (mp *hashmap) delete(key *string) {
	index := hash(key)
	if mp.arr[index].head == nil {
		fmt.Println("\n\nThere is no data to delete..............!!")
	} else {
		mp.listdelete(index, key)
	}
}

// hash funtion
func hash(key *string) int {
	sum := 0
	for _, v := range *key {
		sum += int(v)
	}
	return sum % size
}

func main() {
	var op int
	var key string
	mp := hashmap{}
loop:
	for {
		fmt.Println("\n\n				HASH MAP\n			      -----------")
		fmt.Print("\n\n1--:  INSERT\n\n2--:  DISPLAY\n\n3--:  SEARCH IN MAP\n\n4--:  DELETE FROM MAP\n\n0--:  EXIT\n\nChoose a option   :-")
		fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Print("\nEnter the name  :-")
			fmt.Scan(&key)
			mp.insert(&key)
		case 2:
			mp.display()
		case 3:
			fmt.Print("\nEnter the name to search  :-")
			fmt.Scan(&key)
			fmt.Println(mp.search(&key))
		case 4:
			fmt.Print("\nEnter the name to delete  :-")
			fmt.Scan(&key)
			mp.delete(&key)
		case 0:
			break loop
		default:
			fmt.Println("Please choose correct option............!!")
		}
	}
}
