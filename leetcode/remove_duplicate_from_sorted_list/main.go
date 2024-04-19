package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/solutions/2246708/one-pass-o-1-space-with-go/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// sentinel
	sentinel := &ListNode{Next: head}

	// predecessor = the last node
	// before the sublist of duplicates
	pred := sentinel

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}

			pred.Next = head.Next
		} else {
			pred = pred.Next
		}

		// move forward
		head = head.Next
	}

	return sentinel.Next
}
