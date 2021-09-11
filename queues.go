package main

import (
	"fmt"
	"time"
)

type Queue struct {
	head    *queueNode
	tail    *queueNode
	orderId int
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

func (q *Queue) Dequeue() (string, time.Time) {

	if q.head == nil {
		return "", time.Now()
	}

	value, timeStamp := q.head.value, q.head.timeStamp
	q.head = q.head.next

	return value, timeStamp
}

func (q *Queue) Peek() (string, time.Time) {
	if q.head == nil {
		return "", time.Now()
	}
	return q.head.value, q.head.timeStamp
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

/*
3.2 Stack Min: How would you design a stack which, in addition to push and pop, has a function min
which returns the minimum element? Push, pop and min should all operate in 0(1) time.
*/
/*
type animalQueue struct {
	cat Queue
	dog Queue
}

func (r *minStack) Push(value int) {

	s.s1.Push(value)

	// If S2 is empty or current value is minimum
	if s.s2.IsEmpty() || value <= s.s2.Peek() {
		s.s2.Push(value)
	}

}
func (s *minStack) Pop() int {

	value := s.s1.Pop()

	if value == s.s2.Peek() {
		s.s2.Pop()
	}

	return value
}

func (s *minStack) Peek() int {

	return s.s1.Peek()
}

func (s *minStack) Min() int {

	return s.s2.Peek()
}
*/
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
	name1, time1 := q.Dequeue()
	fmt.Println("Name : ", name1, "Time : ", time1)
	name2, time2 := q.Dequeue()
	fmt.Println("Name : ", name2, "Time : ", time2)

	if time1.Before(time2) {
		fmt.Println("time1 is older ", time1)
	}

	fmt.Println(q.Dequeue())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	fmt.Println(q.IsEmpty())

	/*
		fmt.Println("------------- min Stacks -----------------")

		minSt := minStack{}
		minSt.Push(45)
		minSt.Push(4)
		minSt.Push(3)
		minSt.Push(55)
		minSt.Push(4)
		minSt.Push(2)
		fmt.Println("Peek : ", minSt.Peek())
		fmt.Println("Min : ", minSt.Min())
		fmt.Println(minSt.Pop())
		fmt.Println("Peek : ", minSt.Peek())
		fmt.Println("Min : ", minSt.Min())
	*/
}
