package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
func sortString(str string) string {
	runes := []byte(str)
	sort.Slice(runes, func(a, b int) bool {
		return runes[a] < runes[b]
	})
	return string(runes)
}
*/

func groupAnagrams(strs []string) [][]string {
	hm := make(map[string][]string)
	for _, w := range strs {
		key := sortString(w)
		hm[key] = append(hm[key], w)
	}

	var ans [][]string
	for _, words := range hm {
		ans = append(ans, words)
	}
	return ans
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	fmt.Println("vim-go")
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(str))

}
