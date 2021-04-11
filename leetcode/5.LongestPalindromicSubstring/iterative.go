package main

import "fmt"

func main() {
	text := "hello"
	fmt.Printf("Longest Palindrome of %s: %s\n", text, LongestPalind(text))
	fmt.Printf("Is %s Palindrome? %v\n", text, IsPalind(text, 0, len(text)))
}

func LongestPalind(s string) string {
	start, end := 0, 0

	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			if (j-i) > (end-start) {
				if IsPalindrome(s, i, j) {
					start, end = i, j
				}
			}
		}
	}

	return s[start:end]
}
func IsPalind(s string, start, end int) bool {
	for {
		if end - start <= 1 {return true}
		if s[start] != s[end-1] {return false}
		start++
		end--
	}
}

