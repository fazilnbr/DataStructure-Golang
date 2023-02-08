package trees

import (
	"fmt"
)

type node struct {
	data        int
	left, right *node
}

type BST struct {
	root *node
}

func (Tree *BST) Insert() {
	var data int
	fmt.Println("Enter the data :-  ")
	fmt.Scan(&data)
	newNode := new(node)
	newNode.data = data
	newNode.left = nil
	newNode.right = nil
	temp := Tree.root

	if temp == nil {
		temp = newNode
		Tree.root = newNode

	} else {
		for {
			if data < temp.data {
				if temp.left == nil {
					temp.left = newNode
					break
				} else {
					temp = temp.left
				}
			} else {
				if temp.right == nil {
					temp.right = newNode
					break
				} else {
					temp = temp.right
				}
			}
		}
	}

}

func (Tree *BST) Display() {
	if Tree.root != nil {
		var op int
	loop:
		for {
			fmt.Println("\n\nChoose the format")
			fmt.Print("\n\n1--: InOrder\n\n2--: PostOrder\n\n3--: PreOrder\n\n0--: Go back.\n\n Choose a option :-  ")
			fmt.Scan(&op)
			switch op {
			case 1:
				fmt.Print("INORDER : ")
				inOrder(Tree.root)
			case 2:
				fmt.Print("POSTORDER : ")
				postsOrder(Tree.root)
			case 3:
				fmt.Print("PREORDER : ")
				preOrder(Tree.root)
			case 0:
				break loop
			default:
				fmt.Println("Choose correct option!!!!!!!!!!")

			}
		}
	} else {
		fmt.Println("\n\nThe tree is empty!!!!!!!!!!!!!")
	}
}
func preOrder(root *node) {
	if root != nil {
		fmt.Print(" ", root.data)
		preOrder(root.left)
		preOrder(root.right)

	}
}
func postsOrder(root *node) {
	if root != nil {
		postsOrder(root.left)
		postsOrder(root.right)
		fmt.Print(" ", root.data)

	}
}
func inOrder(root *node) {
	if root != nil {
		inOrder(root.left)
		fmt.Print(" ", root.data)
		inOrder(root.right)

	}
}

func (Tree *BST) Search() bool {
	var num int
	fmt.Print("\nEnter the value :- ")
	fmt.Scan(&num)
	temp := Tree.root
	for temp != nil {
		if num < temp.data {
			temp = temp.left
		} else if num > temp.data {
			temp = temp.right
		} else {
			return true
		}
	}
	return false

}

func (Tree *BST) Delete() {
	var num int
	fmt.Print("\nEnter the number to delete :-")
	fmt.Scan(&num)
	Tree.remove(num, Tree.root, nil)
}

func (Tree *BST) remove(data int, currentNode *node, parentNode *node) {
	for currentNode != nil {
		if data < currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.left
		} else if data > currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.right
		} else {
			if currentNode.left != nil && currentNode.right != nil {
				currentNode.data = getMinimum(currentNode.right)
				Tree.remove(currentNode.data, currentNode.right, currentNode)
			} else {
				if parentNode == nil {
					if currentNode.left == nil {
						Tree.root = currentNode.right
					} else {
						Tree.root = currentNode.left
					}
				} else {
					if parentNode.left == currentNode {
						if parentNode.right == nil {
							parentNode.left = currentNode.left
						} else {
							parentNode.left = currentNode.right
						}
					} else {
						if parentNode.right == nil {
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

func getMinimum(currentNode *node) int {
	if currentNode.left == nil {
		return currentNode.data
	} else {
		return getMinimum(currentNode.left)
	}

}
