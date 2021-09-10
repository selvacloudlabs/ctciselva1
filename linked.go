package main

import (
	"fmt"
)

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

// Insert inserts new node at the end of  from linked list
func insertNode(head *Node, val int) *Node {

	newNode := new(Node)
	newNode.value = val
	newNode.next = nil

	if head == nil {
		head = newNode
		return head
	}

	cur := head

	for cur != nil {

		if cur.next == nil {
			cur.next = newNode
			return head
		}
		cur = cur.next
	}

	return head
}

// Delete inserts new node at the end of  from linked list

func deleteNode(head *Node, val int) *Node {

	if head == nil {
		fmt.Println("Head NULL")
		return nil
	}

	prev := head
	cur := head

	for cur != nil {

		if cur.value == val {
			if head == cur {
				head = cur.next
			} else {
				prev.next = cur.next
			}
		}

		prev = cur
		cur = cur.next
	}

	return head
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

/*
2.4 Partition: Write code to partition a linked list around a value x, such that all nodes less than x come
before all nodes greater than or equal to x. If x is contained within the list the values of x only need
to be after the elements less than x (see below). The partition element x can appear anywhere in the
"right partition"; it does not need to appear between the left and right partitions.
EXAMPLE
Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition= 5]
Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
*/

func partitionList(node *Node, x int) *Node {

	head := node
	tail := node

	printNode(node)

	/* Partition list */
	for node != nil {
		/* Prserve next before changing */
		next := node.next
		if node.value < x {
			/* Insert node at head. */
			node.next = head
			head = node
		} else {
			/* Insert node at tail. */
			tail.next = node
			tail = node
		}
		node = next
	}
	tail.next = nil

	printNode(head)

	return head
}

/*
2.5 Sum Lists: You have two numbers represented by a linked list, where each node contains a single
digit. The digits are stored in reverse order, such that the 1 's digit is at the head of the list. Write a
function that adds the two numbers and returns the sum as a linked list.
EXAMPLE
Input: (7-> 1 -> 6) + (5 -> 9 -> 2).That is,617 + 295.
Output: 2 -> 1 -> 9. That is, 912.
*/
func addLists(l1 *Node, l2 *Node, carry int) *Node {

	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	resultNode := new(Node)
	value := carry

	if l1 != nil {
		value += l1.value
	}
	if l2 != nil {
		value += l2.value
	}
	resultNode.value = value % 10
	if l1 != nil || l2 != nil {
		if value >= 10 {
			resultNode.next = addLists(l1.next, l2.next, 1)
		} else {
			resultNode.next = addLists(l1.next, l2.next, 0)
		}
	}
	return resultNode
}

func main() {

	//head := Node{}
	var listHead *Node

	listHead = insertNode(listHead, 7)
	listHead = insertNode(listHead, 1)
	listHead = insertNode(listHead, 6)
	listHead = insertNode(listHead, 6)
	listHead = insertNode(listHead, 5)
	listHead = insertNode(listHead, 2)
	printNode(listHead)
	//removeDuplicates(listHead)
	//printNode(listHead)
	/*
		listHead = deleteNode(listHead, 9)
		printNode(listHead)
		listHead = deleteNode(listHead, 6)
		printNode(listHead)
		listHead = deleteNode(listHead, 8)
		printNode(listHead)
	*/
	fmt.Println(kthToLast(listHead, 3))

	deleteMiddleNode(kthToLast(listHead, 3))
	printNode(listHead)
	listHead = partitionList(listHead, 5)
	printNode(listHead)

	/* addLists
	var listHeadB *Node
	listHead = insertNode(listHead, 7)
	listHead = insertNode(listHead, 1)
	listHead = insertNode(listHead, 6)
	printNode(listHead)


	listHeadB = insertNode(listHeadB, 5)
	listHeadB = insertNode(listHeadB, 9)
	listHeadB = insertNode(listHeadB, 2)
	printNode(listHeadB)

	printNode(addLists(listHead, listHeadB, 0))
	*/
}
