package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	// Loop through as long nums2 has elements
	for n > 0 {

		// Copy from nums2 if nums1 empty or nums1 is smaller
		if m == 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}

func main() {
	fmt.Println("imerge")

	m := 4
	n := 3
	nums1 := []int{3, 6, 9, 11, 0, 0, 0}
	nums2 := []int{2, 7, 10}

	merge(nums1, m, nums2, n)

	fmt.Println(nums1)

}
