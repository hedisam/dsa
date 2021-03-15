package main

import "fmt"

func main() {
	fmt.Println(shortestPalindrome("enett"))
}

func shortestPalindrome(s string) string {
	return shPal(s, "")
}

func shPal(s string, pref string) string {
	if IsPalindrome(pref+s, 0, len(s)+len(pref)) {
		return pref+s
	}

	i := len(s) - len(pref) - 1
	pref += s[i:i+1]
	return shPal(s, pref)
}

func IsPalindrome(s string, start, end int) bool {
	for {
		if end - start <= 1 {return true}
		if s[start] != s[end-1] {return false}
		start++
		end--
	}
}
