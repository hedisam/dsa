package recursion

func Raise(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * Raise(base, exp-1)
}

func RaiseOpt(base, exp int) int {
	if exp == 0 {
		return 1
	}

	half := RaiseOpt(base, exp/2)
	if exp % 2 == 0 {
		return half * half
	}
	return base * half * half
}
