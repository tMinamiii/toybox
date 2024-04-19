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
func detectCycle(head *ListNode) *ListNode {
	node := head
	reached := make(map[*ListNode]struct{})
	for node != nil { // アドレスを比較
		if _, ok := reached[node]; ok {
			return node
		}
		reached[node] = struct{}{}
		node = node.Next
	}

	return nil
}
