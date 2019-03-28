package main

import (
	"fmt"
	"strings"
)

func LengthOfLongestSubString(s string) int {
	subString := make(map[rune]int)
	start := -1
	maxLength := 0
	for i, c := range []rune(s) {
		if n, ok := subString[c]; ok {
			if n > start {
				start = n
			}
		}
		subString[c] = i
		if i-start > maxLength {
			maxLength = i - start
		}
	}
	return maxLength
}

func main() {
	s := LengthOfLongestSubString("1")
	fmt.Println(s)

	fields := strings.Fields("1  2 3  我 4   5")
	fmt.Println(fields)

	fmt.Println(0xffff)
}
