package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	vals := make(map[int]struct{})
	var prev *ListNode
	current := head
	for current != nil {
		if _, ok := vals[current.Val]; ok {
			prev.Next = current.Next
			current = current.Next
		} else {
			vals[current.Val] = struct{}{}
			prev = current
			current = current.Next
		}
	}
	return head
}
