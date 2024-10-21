package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	currentLen := 0
	s = strings.TrimRight(s, " ")
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			return currentLen
		}
		currentLen++
	}
	return currentLen
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
