package leetcode75

func oddEvenList(head *ListNode) *ListNode {
	evenNodeDummy := &ListNode{Val: 0, Next: nil}
	oddNodeDummy := &ListNode{Val: 0, Next: head}
	evenNode := evenNodeDummy
	last := oddNodeDummy
	oddNode := head
	for oddNode != nil && oddNode.Next != nil {
		evenNode.Next = oddNode.Next
		oddNode.Next = oddNode.Next.Next
		last = oddNode
		oddNode = oddNode.Next
		evenNode = evenNode.Next
	}
	if oddNode != nil && oddNode.Next == nil {
		last = oddNode
	}
	last.Next = evenNodeDummy.Next
	evenNode.Next = nil
	return head
}