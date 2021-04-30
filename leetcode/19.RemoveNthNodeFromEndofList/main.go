package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func list2LinkedList(l []int) *ListNode {
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

func ll2list(head *ListNode) []int {
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	_, head = remKthFromEnd(head, nil, n)
	return head
}

func remKthFromEnd(head, prev *ListNode, n int) (int, *ListNode) {
	if head == nil {
		return 1, nil 
	}

	var i int 
	i, head.Next = remKthFromEnd(head.Next, head, n)

	if i == n {
		head = head.Next
	}

	return i + 1, head 
}