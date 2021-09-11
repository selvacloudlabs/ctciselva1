package main

import (
	"fmt"
)

/*
type Stackable interface {
	Push(int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
}
*/
type Stack struct {
	top *stackNode
}

type stackNode struct {
	value int
	prev  *stackNode
}

func (s *Stack) Push(value int) {
	newNode := new(stackNode)
	newNode.value = value

	if s.top != nil {
		newNode.prev = s.top
	}
	s.top = newNode
}

func (s *Stack) Pop() int {
	if s.top == nil {
		return -1
	}
	ret := s.top.value
	s.top = s.top.prev
	return ret
}

func (s *Stack) Peek() int {
	if s.top == nil {
		return -1
	}
	return s.top.value
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

/*
3.2 Stack Min: How would you design a stack which, in addition to push and pop, has a function min
which returns the minimum element? Push, pop and min should all operate in 0(1) time.
*/

type minStack struct {
	s1 Stack // Main stack
	s2 Stack // To track Min
}

func (s *minStack) Push(value int) {

	s.s1.Push(value)

	/* If S2 is empty or current value is minimum */
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

func main() {
	fmt.Println("------------- Stacks -----------------")

	st := Stack{}

	st.Push(5)
	st.Push(6)
	st.Push(7)
	st.Push(8)
	st.Push(9)
	st.Push(10)

	fmt.Println(st.Peek())
	fmt.Println(st.Peek())
	fmt.Println(st.IsEmpty())

	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())

	fmt.Println(st.IsEmpty())
	fmt.Println(st.Peek())
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
}
