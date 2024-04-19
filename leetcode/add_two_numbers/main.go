package main

import "fmt"

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// sum
	var carryUp int
	resultHead := &ListNode{}
	result := resultHead
	for {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		result.Val = (v1 + v2 + carryUp) % 10
		carryUp = (v1 + v2 + carryUp) / 10
		if l1 == nil && l2 == nil && carryUp == 0 {
			break
		}
		result.Next = &ListNode{}
		result = result.Next
	}

	return resultHead
}

func main() {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	l2 := &ListNode{Val: 9}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Printf("%d, ", result.Val)
		result = result.Next
	}
}
