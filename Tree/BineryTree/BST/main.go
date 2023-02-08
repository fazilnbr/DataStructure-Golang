package main

import (
	"fmt"
	"math"
)

type Node struct {
	data        int
	left, right *Node
}
type BST struct {
	root *Node
}

func main() {
	Tree := BST{}

	fmt.Print("\n 				Binary Search Tree\n 				----------------\n")
loop:
	for {
		fmt.Print("\n\n\n 1-->Insert \n 2-->Condains\n 3-->Remove \n 4-->Display\n 5-->Find Closest \n 6-->LowandHigh\n 0-->Exit\n\n Please choose option : ")
		var op, data int
		fmt.Scan(&op)

		switch op {
		case 1:
			fmt.Print("\nEnter the number :- ")
			fmt.Scan(&data)
			Tree.insert(data)

		case 2:
			fmt.Print("\nEnter the number to find :- ")
			fmt.Scan(&data)
			fmt.Println("\n\n", Tree.Condains(data))
		case 3:
			fmt.Print("\nEnter the number to remove :- ")
			fmt.Scan(&data)
			Tree.remove(data)
		case 4:
			fmt.Print("\nINORDER 	: ")
			Tree.inorder()
			fmt.Print("\nPREORDER 	: ")
			Tree.preorder()
			fmt.Print("\nPOSTORDER 	: ")
			Tree.postorder()
		case 5:
			fmt.Print("\nEnter the number to find closest :- ")
			fmt.Scan(&data)
			fmt.Println("\nClosest value :", Tree.findClosest(data))
		case 6:
			fmt.Println(Tree.rangeSumBST(75, 130))

		case 0:
			break loop
		default:
			fmt.Println("\n Please choose correct option.......!")
		}
	}

}

func (t *BST) insert(data int) {
	currentNode := t.root
	newnode := new(Node)
	newnode.data = data
	if t.root == nil {
		t.root = newnode
		return
	}
	for true {
		if data < currentNode.data {
			if currentNode.left == nil {
				currentNode.left = newnode
				break
			} else {
				currentNode = currentNode.left
			}
		} else {
			if currentNode.right == nil {
				currentNode.right = newnode
				break
			} else {
				currentNode = currentNode.right
			}
		}
	}
}

func (t *BST) Condains(data int) bool {
	if t.root == nil {
		fmt.Println("\n\nThere is no data in tree........! ")
		return false
	}
	currentNode := t.root

	for true {
		if data > currentNode.data {
			if currentNode.right == nil {
				break
			}
			currentNode = currentNode.right
		} else if data < currentNode.data {
			if currentNode.left == nil {
				break
			}
			currentNode = currentNode.left
		} else {
			return true
		}
	}

	return false
}
func (t *BST) remove(data int) {
	t.removeHelper(data, t.root, nil)
}
func (t *BST) removeHelper(data int, currentNode *Node, parentNode *Node) {
	for currentNode != nil {
		if data > currentNode.data {

			parentNode = currentNode
			currentNode = currentNode.right

		} else if data < currentNode.data {

			parentNode = currentNode
			currentNode = currentNode.left

		} else {

			if currentNode.left != nil && currentNode.right != nil {
				currentNode.data = getminvalue(currentNode.right)

				t.removeHelper(currentNode.data, currentNode.right, currentNode)
			} else {
				if parentNode == nil {
					if currentNode.left == nil {
						t.root = currentNode.right
					} else {
						t.root = currentNode.left
					}
				} else {
					if parentNode.left == currentNode {
						if currentNode.right == nil {
							parentNode.left = currentNode.left
						} else {
							parentNode.left = currentNode.right
						}
					} else {
						if currentNode.right == nil {
							parentNode.right = currentNode.left
						} else {
							parentNode.right = currentNode.right
						}
					}
				}
			}
			break

		}
	}

}

func getminvalue(currentNode *Node) int {
	if currentNode.left == nil {
		return currentNode.data
	} else {
		return getminvalue(currentNode.left)
	}
}

func (t *BST) inorder() {
	t.inorderHelper(t.root)
}

func (t *BST) inorderHelper(currentNode *Node) {
	if currentNode != nil {
		t.inorderHelper(currentNode.left)
		fmt.Print(currentNode.data, ", ")
		t.inorderHelper(currentNode.right)
	}
}

func (t *BST) preorder() {
	t.preorderHelper(t.root)
}

func (t *BST) preorderHelper(currentNode *Node) {
	if currentNode != nil {
		fmt.Print(currentNode.data, ", ")
		t.preorderHelper(currentNode.left)
		t.preorderHelper(currentNode.right)
	}
}

func (t *BST) postorder() {
	t.postorderHelper(t.root)
}

func (t *BST) postorderHelper(currentNode *Node) {
	if currentNode != nil {
		t.postorderHelper(currentNode.left)
		t.postorderHelper(currentNode.right)
		fmt.Print(currentNode.data, ", ")
	}
}

func (t *BST) findClosest(target int) int {
	currentNode := t.root
	closest := currentNode.data

	for currentNode != nil {
		if math.Abs(float64(target)-float64(currentNode.data)) > float64(target)-float64(currentNode.data) {
			closest = currentNode.data
		}

		if target > currentNode.data {
			currentNode = currentNode.right
		} else if target < currentNode.data {
			currentNode = currentNode.left
		} else {
			break
		}
	}
	return closest
}

func (t *BST) rangeSumBST(low int, high int) int {
	root := t.root
	return rangeSumBSTHelper(root, low, high)
}

var sum int

func rangeSumBSTHelper(currentNode *Node, low int, high int) int {

	if currentNode != nil {
		sum = rangeSumBSTHelper(currentNode.left, low, high)
		if low <= currentNode.data && high > currentNode.data {
			return sum + currentNode.data
		}
		sum = rangeSumBSTHelper(currentNode.right, low, high)
	}
	return sum
}
