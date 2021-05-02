package main 

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    _, ok := isPalind(head, head.Next)
	return ok 
}

func isPalind(head, tail *ListNode) (*ListNode, bool) {
	if tail == nil {
		return nil, true 
	}

	headNext, subPalindrome := isPalind(head, tail.Next)
	if !subPalindrome {return nil, false}
	
	if headNext == nil {
		headNext = head 
	}

	return headNext.Next, headNext.Val == tail.Val
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