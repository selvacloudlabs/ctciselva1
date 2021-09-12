package main

import "fmt"

func swapOddEvenBits(x int) int {
	return (((x & 0xaaaaaaaa) >> 1) | ((x & 0x55555555) << 1))
}

func countNumBits(n int) int {
	count := 0

	for n != 0 {
		fmt.Printf(" n = %d [%b] (n-1) = %d [%b]  n&(n-1) = %d [%b] \n", n, n, n-1, n-1, n&(n-1), n&(n-1))
		count++
		n &= (n - 1)
	}
	return count
}

func main() {

	x := 5
	fmt.Printf("%x\n", swapOddEvenBits(x))

	n := 74

	//fmt.Printf(" n = %d (n-1) = %d  n*(n-1) = %d \n", n, n-1, n*(n-1))
	//fmt.Printf(" n = %d [%b] (n-1) = %d [%b]  n&(n-1) = %d [%b] \n", n, n, n-1, n-1, n&(n-1), n&(n-1))

	fmt.Println(countNumBits(n))
}
