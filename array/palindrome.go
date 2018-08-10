package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkPalindrome("arara"))
	fmt.Println(checkPalindrome("this isn't a palindrome"))
}

func checkPalindrome(palavra string) bool {
	s := []rune(strings.ToLower(palavra))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Print(s[i], " e ", s[j], "\n")
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
