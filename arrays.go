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

func main() {
	fmt.Println("################# 1.1 Unique string ################")
	
	fmt.Println(isUniqueChars("Hello"))
	fmt.Println(isUniqueChars("Helo"))
	fmt.Println("################# 1.1 Unique string ################")
	fmt.Println(arePermutation("Hello","ollHe"))
	fmt.Println(arePermutation("Hello","oloHe"))

}

