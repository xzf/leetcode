package main

import "fmt"

//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	bSlice := []byte(s)
	var max int
	for index, c := range bSlice {
		if index == len(bSlice)-1 {
			break
		}
		subSlice := []byte{c}
		for i := index + 1; i < len(bSlice); i++ {
			for _, oc := range subSlice {
				if bSlice[i] == oc {
					break
				}
			}
			subSlice = append(subSlice, bSlice[i])
			fmt.Println(string(subSlice))

		}
		if len(subSlice) > max {
			max = len(subSlice)
		}
	}
	return max
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
