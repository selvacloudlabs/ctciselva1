package main

import (
	"fmt"
)

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

var listHead *Node

// Insert inserts new node at the end of  from linked list
func insertNode(head *Node, val int) {

	newNode := Node{}
	newNode.value = val
	newNode.next = nil

	if head == nil {
		listHead = &newNode
		return
	}

	cur := head

	for cur != nil {

		if cur.next == nil {
			cur.next = &newNode
			return
		}
		cur = cur.next
	}

	return
}

// Delete inserts new node at the end of  from linked list
func deleteNode(head *Node, val int) {

	if head == nil {
		fmt.Println("Head NULL")
		return
	}

	prev := head
	cur := head

	for cur != nil {

		if cur.value == val {
			if listHead == cur {
				listHead = cur.next
			} else {
				prev.next = cur.next
			}
		}

		prev = cur
		cur = cur.next
	}

	return
}

// Insert inserts new node at the end of  from linked list
func printNode(head *Node) {

	if head == nil {
		fmt.Println("Head NULL")
		return
	}
	fmt.Println("------------------ Printing LIst ----------------")

	cur := head

	for cur != nil {
		fmt.Println("Node : ", cur)
		cur = cur.next
	}

	return
}

// 2.1 Remove Dups: Write code to remove duplicates from an unsorted linked list.
// Method 2
// space complexity - O(n)
// time complexity - O(n)
/**
 * [removeDuplicates1 - Remove duplicates from the list using hash table]
 * @param head [head of list]
 */

func removeDuplicates(head *Node) {
	//if ( head == nil ||  (head && (head.next == nil ))) {
	if head == nil {
		return
	}
	foundHashTable := make(map[int]bool)
	prev := head
	cur := head.next

	foundHashTable[head.value] = true

	for cur != nil {
		if foundHashTable[cur.value] {
			prev.next = cur.next
		} else {
			foundHashTable[cur.value] = true
			prev = cur
		}
		cur = cur.next
	}
}

/*
2.2 Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
*/
func kthToLast(head *Node, k int) *Node {
	p1 := head
	p2 := head

	/* Move p1 k nodes into the list.*/
	for i := 0; i < k; i++ {
		if p1 == nil {
			return nil
		} // Out of bounds
		p1 = p1.next
	}

	/* Move them at the same pace. When p1 hits the end,
	 * p2 will be at the right element. */
	for p1 != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2
}

/*
2.3 Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but
the first and last node, not necessarily the exact middle) of a singly linked list, given only access to
that node.
EXAMPLE
lnput:the node c from the linked list a->b->c->d->e->f
Result: nothing is returned, but the new linked list looks like a ->b->d- >e- >f
*/
func deleteMiddleNode(midNode *Node) bool {

	if midNode == nil || midNode.next == nil {
		return false // Failure
	}
	nextNode := midNode.next
	midNode.value = nextNode.value
	midNode.next = nextNode.next
	return true
}

func main() {

	//head := Node{}

	//head.value = 4
	//printNode(listHead)
	insertNode(listHead, 4)

	printNode(listHead)

	insertNode(listHead, 6)
	insertNode(listHead, 6)
	insertNode(listHead, 8)
	insertNode(listHead, 9)
	insertNode(listHead, 13)
	insertNode(listHead, 14)
	insertNode(listHead, 25)
	printNode(listHead)
	removeDuplicates(listHead)
	printNode(listHead)
	/*
		deleteNode(listHead, 4)
		printNode(listHead)
		deleteNode(listHead, 6)
		printNode(listHead)
		deleteNode(listHead, 8)
		printNode(listHead)
	*/
	fmt.Println(kthToLast(listHead, 3))

	deleteMiddleNode(kthToLast(listHead, 3))
	printNode(listHead)

}
