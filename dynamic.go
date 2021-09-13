package main

import "fmt"

// Swap the i-th byte and j-th byte of the string
func swap(s string, i, j int) string {
	var result []byte
	for k := 0; k < len(s); k++ {
		if k == i {
			result = append(result, s[j])
		} else if k == j {
			result = append(result, s[i])
		} else {
			result = append(result, s[k])
		}
	}
	return string(result)
}

// Function to find all Permutations of a given string str[i:n]
// containing all distinct characters
func permutations(str string, i, n int) {
	// base condition
	if i == n-1 {
		fmt.Println(str)
		return
	}

	// process each character of the remaining string
	for j := i; j < n; j++ {
		// swap character at index i with current character
		str = swap(str, i, j)

		// recursion for string [i+1:n]
		permutations(str, i+1, n)

		// backtrack (restore the string to its original state)
		str = swap(str, i, j)
	}
}

func main() {
	str := "ABC"
	permutations(str, 0, len(str))
}
