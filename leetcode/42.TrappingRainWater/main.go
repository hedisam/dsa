package main

import "fmt"

func main() {
	h := []int{0,1,2,0,3,0,1,2,0,0,4,2,1,2,5,0,1,2,0,2}
	fmt.Println(Trap(h))
}

func Trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	// find the highest bar
	var highest int
	for i, h := range height {
		if h >= height[highest] {
			highest = i
		}
	}

	// we compute the trapped water by dividing the bars into two sections
	// first section consists of the bars to the left of the highest bar
	// the second one consists of the bars to the right of the highest bar
	// we use a reverse version of the algorithm for the second section

	var total int

	// doing the left side
	water := trapRec(height, 0, 1, 0, highest)
	total += water

	// doing the right side
	water = trapRecRev(height, len(height)-1, len(height)-2, 0, highest)
	total += water

	return total
}

func trapRecRev(height []int, start, end, blocksBetween, highest int) int {
	if end < highest {
		return 0
	}

	var nextStart, nextEnd int
	var nextBlocksBetween int
	var totalWater int

	if height[start] <= height[end] {
		totalWater = (start - end - 1) * height[start] - blocksBetween
		nextStart = end
		nextEnd = nextStart-1
		nextBlocksBetween = 0
		goto recur
	}

	nextStart = start
	nextEnd = end - 1
	nextBlocksBetween = blocksBetween + height[end]

recur:
	trapped := trapRecRev(height, nextStart, nextEnd, nextBlocksBetween, highest)
	totalWater += trapped
	return totalWater
}

func trapRec(height []int, start, end int, blocksBetween int, highest int) int {
	if end > highest {
		return 0
	}

	var nextStart, nextEnd int
	var nextBlocksBetween int
	var totalWater int

	if height[start] <= height[end] {
		totalWater = (end - start - 1) * height[start] - blocksBetween
		nextStart = end
		nextEnd = nextStart+1
		nextBlocksBetween = 0
		goto recur
	}

	nextStart = start
	nextEnd = end + 1
	nextBlocksBetween = blocksBetween + height[end]

recur:
	trapped := trapRec(height, nextStart, nextEnd, nextBlocksBetween, highest)
	totalWater += trapped
	return totalWater
}
