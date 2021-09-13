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

var hashMap map[int]*Node

func getNode(node *Node, data int) {

	if node == nil {
		return
	}

	if node.data == data {
		fmt.Println(" getNode : Match found : ", node)
		hashMap[data] = node
	}

	getNode(node.left, data)
	getNode(node.right, data)

	return
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

/*
4.8 First Common Ancestor: Design an algorithm and write code to find the first common ancestor
of two nodes in a binary tree. Avoid storing additional nodes in a data structure. NOTE: This is not
necessarily a binary search tree.
*/

func findLowestCommonAncestor(currentNode *Node, p *Node, q *Node) bool {

	// If reached the end of a branch, return false.
	if currentNode == nil {
		return false
	}

	// Left Recursion. If left recursion returns true, set left = 1 else 0
	left := findLowestCommonAncestor(currentNode.left, p, q)

	// Right Recursion
	right := findLowestCommonAncestor(currentNode.right, p, q)

	// If the current node is one of p or q
	mid := (currentNode == p || currentNode == q)

	// If any two of the flags left, right or mid become True
	if (left && mid) || (mid && right) || (left && right) {
		fmt.Println("The Common Ancestor is ", currentNode)
	}

	// Return true if any one of the three bool values is True.
	return (mid || left || right)
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
	//var res *Node
	hashMap = make(map[int]*Node)
	getNode(root, 1)
	getNode(root, 2)
	getNode(root, 3)
	getNode(root, 4)
	getNode(root, 5)
	getNode(root, 6)
	getNode(root, 7)
	fmt.Println("Main match : ", hashMap)
	//fmt.Println("Main match : ", getNode(root, 6))

	findLowestCommonAncestor(root, hashMap[2], hashMap[6])
}
