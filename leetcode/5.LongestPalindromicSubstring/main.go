package main


type Point struct {
	Start int
	End int
}

func PointLength(p *Point) int {
	return p.End - p.Start
}

func LongestPalindrome(s string, start, end int, memo map[Point]Point) Point {
	if value, ok := memo[Point{Start: start, End: end}]; ok {
		return value
	}
	if IsPalindrome(s, start, end) {
		p := Point{Start: start, End: end}
		memo[p] = p
		return p
	}

	left := LongestPalindrome(s, start, end-1, memo)
	right := LongestPalindrome(s, start+1, end, memo)

	if PointLength(&left) > PointLength(&right) {
		memo[Point{Start: start, End: end}] = left
		return left
	}
	memo[Point{Start: start, End: end}] = right
	return right
}

func IsPalindrome(s string, start, end int) bool {
	if end-start <= 1 {return true}

	return s[start] == s[end-1] && IsPalindrome(s, start+1, end-1)
}
