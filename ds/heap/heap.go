package heap


func right(i int) int {
	return 2 * i + 2
}

func left(i int) int {
	return 2 * i + 1
}

func parent(i int) int {
	return (i - 1) / 2
}