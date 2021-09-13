package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/* 1.1 Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
cannot use additional data structures?
*/
func isUniqueChars(str string) bool {

	var charTable [128]bool

	if len(str) > 128 {
		return false
	}

	for idx := 0; idx < len(str); idx++ {

		key := str[idx]

		if charTable[key] {
			return false
		}
		charTable[key] = true
	}
	return true
}

func arePermutation(str1 string, str2 string) bool {

	var freqTable [128]int

	if len(str1) != len(str2) {
		return false
	}

	for idx := 0; idx < len(str1); idx++ {

		key := str1[idx]

		freqTable[key]++
	}

	for idx := 0; idx < len(str2); idx++ {

		key := str2[idx]

		freqTable[key]--

		if freqTable[key] < 0 {
			return false
		}
	}
	return true
}

// O(n) time with O(n) extra space.
func URLify(input string, truelen int) string {
	spaces := 0

	for idx_in := 0; idx_in < truelen; idx_in++ {

		if input[idx_in] == byte(' ') {
			spaces++
		}
	}
	// +2 for each space for the extra "20"s.
	output := make([]byte, truelen+(2*spaces))
	idx_out := truelen + spaces*2 - 1

	for idx_in := truelen - 1; idx_in >= 0; idx_in-- {
		if input[idx_in] == byte(' ') {
			output[idx_out] = byte('0')
			output[idx_out-1] = byte('2')
			output[idx_out-2] = byte('%')
			idx_out -= 3
		} else {
			output[idx_out] = input[idx_in]
			idx_out--
		}
	}
	return string(output)
}

/*
 * Given a string, write a function to check if it is a permutation of a pallindrome.
 *
 * Solution Philosophy:
 * For a string to be pallindrome, it should be able to spelled backward and forward the same.
 * Therefore the chars in string should fit one of the two possibilities:
 *  - Each char appear even number of times in the string ( even length string )
 *  - Each char should appear even number of times, except just one char ( odd length string )
 *
 * We won't care about the case of the letter
 */

func getFreqTableKey(ch byte) int {

	key := -1

	if ch >= 'a' && ch <= 'z' {
		key = int(ch - 'a')
	} else if ch >= 'A' && ch <= 'Z' {
		key = int(ch - 'A')
	}

	return key
}

func isPermutationOfPallindrome(str string) bool {

	var freqTable [26]int
	oddCount := 0
	key := -1

	for idx := 0; idx < len(str); idx++ {

		key = getFreqTableKey(str[idx])

		if key != -1 {

			freqTable[key]++

			if freqTable[key]%2 == 1 {
				oddCount++
			} else {
				oddCount--
			}
		}

	}
	return (oddCount <= 1)
}

/*
1.6 String Compression: Implement a method to perform basic string compression using the counts
of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
"compressed" string would not become smaller than the original string, your method should return
the original string. You can assume the string has only uppercase and lowercase letters (a - z).
*/
func compressString(origStr string) string {

	origLen := len(origStr)
	var compStr string

	if origLen < 2 {
		return origStr
	}

	count := 1
	for idx := 1; idx < origLen; idx++ {

		if origStr[idx-1] == origStr[idx] {
			count++
		} else {
			compStr += string(origStr[idx-1]) + strconv.Itoa(count)
			count = 1
		}
		if len(compStr) >= origLen {
			return origStr
		}
	}
	compStr += string(origStr[origLen-1]) + strconv.Itoa(count)

	if len(compStr) >= origLen {
		return origStr
	}
	return compStr
}

/*1.9 String Rotation: Assume you have a method i 5Su b 5 tr ing which checks if one word is a substring
of another. Given two strings, 51 and 52, write code to check if 52 is a rotation of 51 using only one
call to i5Sub5tring (e.g., "waterbottle" is a rotation of" erbottlewat").
*/

func isRotation(s1 string, s2 string) bool {

	/* check that s1 and s2 are equal length and not empty */
	if len(s1) == len(s2) && len(s1) > 0 {

		/* concatenate s1 and s1 within new buffer */
		s1s1 := s1 + s1
		return strings.Contains(s1s1, s2)
	}
	return false
}

/*
 * Problem: There are three possible edits that can be performed on a string.
 * 1. Insert a char.
 * 2. Delete a char.
 * 3. Replace a char.
 *
 * Given two strings, determine if they are one or 0 edit away.
 *
 * Approach :
 * 1. Case when strings are of some length --> possible edit is replace.
 *    If there are more than one mismatch, return false
 *
 * 2. Case when One string is bigger than another
 *    Smaller string ------------> Bigger String
 *                     insert
 *                     delete
 *    smaller string <-----------  Bigger String
 *
 *    Idea is check if there are more than one mismatch discounting the already
 *    difference in the string. Therefore for first mismatch we do not move the pointer
 *    pointing to smaller string, and then expect it to match from next char of bigger
 *    string.
 */

func oneEditAway(str1 string, str2 string) bool {

	len1 := len(str1)
	len2 := len(str2)
	var smaller, bigger string

	if math.Abs(float64(len1-len2)) > 1 {
		return false
	}

	if len1 < len2 {
		smaller = str1
		bigger = str2
	} else {
		smaller = str2
		bigger = str1

	}

	var i, j int
	mismatchDone := false

	for i < len(smaller) && j < len(bigger) {

		if smaller[i] != bigger[j] {
			if mismatchDone {
				return false
			}
			mismatchDone = true
			if len1 == len2 {
				i++ //case of replace
			} // else case of replace.dont move small pointer
		} else {
			i++ //move short pointer if its a match, dont move it in case of first mismatch
		}
		j++ //always move long string pointer.
	}
	return true
}

/**
 * Cracking the coding interview 1.8
 * Write a space efficient algorithm, such that if an element in MxN is 0, the entire row and column containing it are 0.
 *
 * Approach:
 * We can use a boolean matrix of MxN or a bit vector to mark row and columns to be nullified in first iteration, but it wont be space efficient.
 * More space efficient would be to first check first row and column and if any of them contains zero, mark them to be nullified using two boolearn vars
 * let's say firstRow and firstCol, and then iterate through rest of the matrix and store information in first row column elements, only when that row
 * and column is to be marked for nullified, this way we will only change values in first row and column which are already going to be 0 in final solution.
 */

func nullifyRow(matrix [][]int, numColumns int, rowId int) {
	for j := 0; j < numColumns; j++ {
		matrix[rowId][j] = 0
	}
}

func nullifyColumn(matrix [][]int, numRows int, colId int) {
	for i := 0; i < numRows; i++ {
		matrix[i][colId] = 0
	}
}

func nullifyMatrix(matrix [][]int) {

	numRows := len(matrix)
	numColumns := len(matrix[0])
	nullifyFirstRow := false

	//first row
	for i := 0; i < numColumns; i++ {
		if matrix[0][i] == 0 {
			nullifyFirstRow = true
			break
		}
	}

	//rest of the matrix
	for i := 1; i < numRows; i++ {
		nullifyThisRow := false
		for j := 0; j < numColumns; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				nullifyThisRow = true
			}
		}
		if nullifyThisRow {
			nullifyRow(matrix, numColumns, i)
		}
	}

	//now we know which column to be nullify using information stored in previous step.
	//cols first
	for j := 0; j < numColumns; j++ {
		if matrix[0][j] == 0 {
			nullifyColumn(matrix, numRows, j)
		}
	}

	//now first row
	if nullifyFirstRow {
		nullifyRow(matrix, numColumns, 0)
	}

}

/*
 * Cracking the coding interview edition 6
 * Problem 1.7 Rotate a matrix by 90' clockwise ( or anticlockwise).
 * Solution : I have done it two ways.
 * Approach 1: Take transpose of matrix and then reverse the rows for clockwise 90' rotation.
 * 			   Obviously if we reverse the columns we will get anticlockwise 90' rotation.
 * Approach 2: As mentioned in the book, rotating invididual elements layer by layer.
 * 			   I have solved it perform anticlockwise 90' rotation, it can be done similarly for clockwise rotatation.
 */

func transposeMatrix(matrix [][]int, N int) {

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if i != j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
}

func reverseRow(row []int, N int) {
	for i := 0; i < N/2; i++ {
		row[i], row[N-i-1] = row[N-i-1], row[i]
	}
}

func rotateMatrix(matrix [][]int) bool {

	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return false // Not a square
	}
	N := len(matrix)

	//transpose matrix
	transposeMatrix(matrix, N)

	// reverse each row
	for i := 0; i < N; i++ {
		reverseRow(matrix[i], N)
	}
	return true
}

func rotateMatrix2(matrix [][]int) bool {

	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return false // Not a square
	}
	n := len(matrix)

	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		//fmt.Printf("\nOUTER : layer = %d ; first = %d ; last = %d\n",layer,first,last)
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i] // save top
			//fmt.Printf("\n    INNER : i = %d ; offset = %d ; last-offset = %d\n",i,offset,last-offset)

			//fmt.Printf("        BFR : matrix[first][i] = %d  matrix[last-offset][first] = %d matrix[last][last - offset] = %d matrix[i][last] = %d\n",matrix[first][i], matrix[last-offset][first],matrix[last][last - offset],matrix[i][last])
			// left -> top
			matrix[first][i] = matrix[last-offset][first]

			// bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]

			// right -> bottom
			matrix[last][last-offset] = matrix[i][last]

			// top -> right
			matrix[i][last] = top // right <- saved top
			//fmt.Printf("        AFR : matrix[first][i] = %d  matrix[last-offset][first] = %d matrix[last][last - offset] = %d matrix[i][last] = %d\n",matrix[first][i], matrix[last-offset][first],matrix[last][last - offset],matrix[i][last])
		}
	}
	return true
}

func main() {
	fmt.Println("################# 1.1 Unique string ################")

	fmt.Println(isUniqueChars("Hello"))
	fmt.Println(isUniqueChars("Helo"))
	fmt.Println("################# 1.2 Are Palindrome ################")
	fmt.Println(arePermutation("Hello", "ollHe"))
	fmt.Println(arePermutation("Hello", "oloHe"))
	fmt.Println("################# 1.3 URLify ################")
	fmt.Println(URLify("H el lo      ", 7))
	fmt.Println("################# 1.4 Permutation palindrome ################")
	fmt.Println(isPermutationOfPallindrome("TactCooa"))

	fmt.Println("################# 1.5 One Edit away ################")
	fmt.Println(oneEditAway("pal", "pa3"))

	fmt.Println("################# 1.6 Compress STring ################")
	fmt.Println(compressString("aaaabbb"))

	fmt.Println("################# 1.7 Rotate Matrix ################")
	matrix := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}

	fmt.Println(matrix)
	rotateMatrix(matrix)
	fmt.Println(matrix)

	/*
	   	fmt.Println("################# 1.8 Zero Matrix ################")
	           matrix1 := [][]int{
	   		        []int{1,2,0},
	   		        []int{4,0,6},
	   		        []int{7,8,9},
	                         }


	   	fmt.Println(matrix1)
	   	nullifyMatrix(matrix1)
	   	fmt.Println(matrix1)


	   	fmt.Println("################# 1.9 String Rotation ################")
	   	fmt.Println(isRotation("waterbottle","erbottlewatq"))
	*/

}
