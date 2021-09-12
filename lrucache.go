package main

import (
	"fmt"
)

type Node struct {
	key   int
	value string
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *LinkedList) InsertAtHead(key int, value string) {

	newNode := &Node{key: key, value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.size++
}

func (l *LinkedList) InsertAtTail(key int, value string) {
	newNode := &Node{key: key, value: value}
	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
		newNode.next = nil
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		newNode.next = nil
	}
	l.size++
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}

func (l *LinkedList) GetTail() *Node {
	return l.tail
}

func (l *LinkedList) RemoveNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == l.head {

		l.head = l.head.next
	}
	if node == l.tail {

		l.tail = l.tail.prev
		l.size--
	}
	return node
}

func (l *LinkedList) RemoveHead() *Node {
	return l.RemoveNode(l.head)
}

type LRUCache struct {
	capacity int

	//LinkedListNode holds key and value pairs
	cacheMap  map[int]*Node
	cacheList LinkedList
}

func (lr *LRUCache) Get(key int) *Node {

	hashNode, ok := lr.cacheMap[key]

	if !ok {
		return nil
	} else {
		lr.cacheList.RemoveNode(hashNode)
		lr.cacheList.InsertAtTail(key, hashNode.value)
		return lr.cacheList.GetTail()
	}
}

func (lr *LRUCache) Set(key int, value string) {
	if _, ok := lr.cacheMap[key]; !ok {
		if lr.cacheList.size >= lr.capacity {
			delete(lr.cacheMap, lr.cacheList.GetHead().key)
			lr.cacheList.RemoveHead()
		}
		lr.cacheList.InsertAtTail(key, value)
		lr.cacheMap[key] = lr.cacheList.GetTail()

	} else {
		lr.cacheList.RemoveNode(lr.cacheMap[key])
		lr.cacheList.InsertAtTail(key, value)
		lr.cacheMap[key] = lr.cacheList.GetTail()
	}
}

func (lr *LRUCache) Print() {
	curr := lr.cacheList.head
	for curr != nil {
		fmt.Printf("(%v,%s)", curr.key, curr.value)
		curr = curr.next
	}
	fmt.Println("")
}

func main() {
	cache := &LRUCache{capacity: 3, cacheMap: make(map[int]*Node)}
	fmt.Println("The most recently used items are: (key, value)")
	cache.Set(10, "Food")
	cache.Print()

	cache.Set(15, "Chair")
	cache.Print()

	cache.Set(20, "Clothes")
	cache.Print()

	cache.Set(25, "Tables")
	cache.Print()

	cache.Set(5, "Sofa")
	cache.Print()

	cache.Get(25)
	cache.Print()
}
