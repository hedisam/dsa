package recursion

func IsPalindrome(s string) bool {
	length := len(s)
	if length <= 1 {
		return true
	}
	return s[0] == s[length-1] && IsPalindrome(s[1:length-1])
}
