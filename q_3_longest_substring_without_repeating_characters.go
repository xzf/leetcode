package main

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var max int
	for index := 0; index < len(s)-1; index++ {
		var endIndex int
		for i := index + 1; i < len(s); i++ {
			this := string(s[i])
			if strings.Contains(s[index:i], this) {
				break
			}
			endIndex = i
		}
		var sub string
		if endIndex >= len(s) {
			sub = s[index:]
		} else if endIndex == 0 {
			sub = s[index:index+1]
		}else {
			sub = s[index:endIndex+1]
		}
		if len(sub) > max {
			max = len(sub)
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
