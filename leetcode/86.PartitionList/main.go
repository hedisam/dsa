package main

import "fmt"

func main() {
	head := array2LinkedList([]int{1, 4, 3, 2, 5, 2})
	parted := partition(head, 3)

	arr := ll2array(parted)
	fmt.Println(arr)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	left, right, leftTail := part(head, x)
	if leftTail == nil {
		return right
	}
	leftTail.Next = right
	return left
}

func part(org *ListNode, x int) (*ListNode, *ListNode, *ListNode) {
	if org == nil {
		return nil, nil, nil
	}

	if org.Val < x {
		left := &ListNode{Val: org.Val}
		leftNext, right, leftTail := part(org.Next, x)
		left.Next = leftNext
		if leftTail == nil {
			leftTail = left
		}
		return left, right, leftTail
	} else {
		right := &ListNode{Val: org.Val}
		left, rightNext, leftTail := part(org.Next, x)
		right.Next = rightNext
		if leftTail == nil {
			leftTail = left
		}
		return left, right, leftTail
	}
}

func array2LinkedList(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}

	head := &ListNode{Val: l[0]}

	var node *ListNode = head
	for i := 1; i < len(l); i++ {
		node.Next = &ListNode{Val: l[i]}
		node = node.Next
	}

	return head
}

func ll2array(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return arr
}