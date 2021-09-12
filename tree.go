package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func inOrder(node *Node) {

	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Println(node.data)
	inOrder(node.right)
}
func preOrder(node *Node) {

	if node == nil {
		return
	}

	fmt.Println(node.data)
	preOrder(node.left)
	preOrder(node.right)
}
func postOrder(node *Node) {

	if node == nil {
		return
	}

	postOrder(node.left)
	postOrder(node.right)
	fmt.Println(node.data)
}

func createBST(arr []int, low int, high int) *Node {

	if low > high {
		return nil
	}
	mid := (low + high) / 2
	node := new(Node)
	node.data = arr[mid]
	node.left = createBST(arr, low, mid-1)
	node.right = createBST(arr, mid+1, high)

	return node
}

func createMinimalBST(arr []int) *Node {

	if len(arr) == 0 {
		return nil
	}

	return createBST(arr, 0, len(arr)-1)
}

/*
4.5 Validate BST: Implement a function to check if a binary tree is a binary search tree.
*/

func validateBST(root *Node, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.data < min || root.data > max {
		return false
	}
	fmt.Println(root.data, min, max)
	return validateBST(root.left, min, root.data-1) && validateBST(root.right, root.data+1, max)
}

func main() {

	A := []int{1, 2, 3, 4, 5, 6, 7}

	/*
	   Convert List to BST {1,2,3,4,5,6,7}
	                    4
	                 /     \
	               2         6
	             /   \     /   \
	            1     3   5     7
	*/
	root := createMinimalBST(A)

	fmt.Println("InOrder Traversal of constructed BST :")
	inOrder(root)
	fmt.Println("PreOrder Traversal of constructed BST :")
	preOrder(root)
	fmt.Println("PostOrder Traversal of constructed BST :")
	postOrder(root)

	fmt.Println("Validating if BST :")
	fmt.Println(validateBST(root, -2e9, 2e9))
}
