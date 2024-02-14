package leetcode75

func ReverseList(head *ListNode) *ListNode {
	var dummy *ListNode
	node := head
	last := dummy
	for node != nil {
		tmp := node.Next
		node.Next = last
		last = node
		node = tmp
	}
	return last
}