package main

import "fmt"

/* 1.1 Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
cannot use additional data structures?
*/
func isUniqueChars(str string) bool {

	if (len(str) > 128){
		return false;
	}

	var char_set [128]bool;

	for i := 0; i < len(str); i++{

	    val := str[i];

	    if (char_set[val]){
		return false;
	    }
	    char_set[val] = true;
	}
	return true;
}


func arePermutation(str1 string , str2 string) bool {

	if(len(str1) != len(str2)) { 
    	   return false;
        }
	var count [128]int;

  	for i := 0; i < len(str1); i++ {

            val := str1[i];

            count[val]++;
  	}

	for i := 0; i < len(str2); i++ {

	    val := str2[i];
            count[val]--;

            if count[val]<0 { 
              return false;
            }
        }
        return true;
}

// O(n) time with O(n) extra space.
func URLify(input string, truelen int) string {
	spaces := 0

	for idx_in := 0; idx_in< truelen; idx_in++ {
	
		if input[idx_in] == byte(' ') {
			spaces++
		}
	}
	// +2 for each space for the extra "20"s.
	output := make([]byte, truelen+(2*spaces))
	idx_out := truelen + spaces * 2 - 1;

	for idx_in := truelen - 1; idx_in>=0 ;  idx_in-- {
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



func main() {
	fmt.Println("################# 1.1 Unique string ################")
	
	fmt.Println(isUniqueChars("Hello"))
	fmt.Println(isUniqueChars("Helo"))
	fmt.Println("################# 1.2 Are Palindrome ################")
	fmt.Println(arePermutation("Hello","ollHe"))
	fmt.Println(arePermutation("Hello","oloHe"))
	fmt.Println("################# 1.3 URLify ################")
	fmt.Println(URLify("H el lo      ",7))


}

