package main

import (
	"fmt"
	"time"
)

type Queue struct {
	head *queueNode
	tail *queueNode
}

type queueNode struct {
	value     string
	timeStamp time.Time
	next      *queueNode
}

func (q *Queue) Enqueue(value string) {
	newNode := new(queueNode)
	newNode.value = value
	newNode.timeStamp = time.Now()

	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue) Dequeue() *queueNode {

	if q.head == nil {
		return nil
	}

	ret := q.head
	q.head = q.head.next

	return ret
}

func (q *Queue) Peek() *queueNode {

	return q.head
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

/*
3.6 Animal Shelter: An animal shelter, which holds only dogs and cats, operates on a strictly"first in, first
out" basis. People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of
that type). They cannot select which specific animal they would like. Create the data structures to
maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog,
and dequeueCat. You may use the built-in Linkedlist data structure.
*/
type animalQueue struct {
	cat Queue
	dog Queue
}

func (aq *animalQueue) EnqueueAnimal(animalType string, animalName string) {

	if animalType == "cat" {
		aq.cat.Enqueue(animalName)
	} else if animalType == "dog" {
		aq.dog.Enqueue(animalName)
	} else {
		fmt.Println("Invalid animal Type")
	}
}

func (aq *animalQueue) DequeueAnimal() *queueNode {

	c := aq.cat.Peek()
	d := aq.dog.Peek()

	if c.timeStamp.Before(d.timeStamp) {
		return aq.cat.Dequeue()
	} else {
		return aq.dog.Dequeue()
	}
}

func (aq *animalQueue) DequeueCat() *queueNode {

	return aq.cat.Dequeue()
}

func (aq *animalQueue) DequeueDog() *queueNode {

	return aq.dog.Dequeue()
}

func main() {
	fmt.Println("------------- Queue -----------------")

	q := Queue{}

	q.Enqueue("catA")
	q.Enqueue("catB")
	q.Enqueue("dogA")
	q.Enqueue("dogB")
	q.Enqueue("dogC")
	q.Enqueue("catC")

	fmt.Println(q.Peek())
	//qn1 := q.Dequeue()
	fmt.Println(q.Dequeue())
	//qn2 := q.Dequeue()
	//fmt.Println("Name : ", qn2.value, "Time : ", qn2.timeStamp)

	//if qn1.timeStamp.Before(qn2.timeStamp) {
	//	fmt.Println("time1 is older ", qn1.timeStamp)
	//}

	fmt.Println(q.Dequeue())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	fmt.Println(q.IsEmpty())

	fmt.Println("------------- ANimal Queue -----------------")
	aq := animalQueue{}

	aq.EnqueueAnimal("cat", "catA")
	aq.EnqueueAnimal("dog", "dogA")
	aq.EnqueueAnimal("cat", "catB")
	aq.EnqueueAnimal("cat", "catC")
	aq.EnqueueAnimal("dog", "dogB")
	aq.EnqueueAnimal("dog", "dogC")
	aq.EnqueueAnimal("cat", "catD")

	fmt.Println(aq.DequeueAnimal())
	fmt.Println(aq.DequeueCat())
	fmt.Println(aq.DequeueCat())
	fmt.Println(aq.DequeueDog())
	fmt.Println(aq.DequeueDog())

}
