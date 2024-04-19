package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// slowは1つずつ進め、fastは2つずつ進める
	// 差が広がっていくので、循環しているといずれ同じノードを指す
	slow := head       // 1つずつ進める
	fast := head.Next  // 2つずつ進める
	for slow != fast { // アドレスを比較
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
