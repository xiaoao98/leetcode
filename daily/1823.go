package daily

func findTheWinner(n int, k int) int { //or use queue
	if k == 1 {
		return n
	}
	head := &ListNode{Val: 1}
	node := head
	for i := 2; i <= n; i++ {
		newNode := &ListNode{Val: i}
		node.Next = newNode
		node = node.Next
	}
	node.Next = head
	count := n
	node = head
	for count > 1 {
		for i := 2; i < k; i++ {
			node = node.Next
		}
		node.Next = node.Next.Next
		node = node.Next
		count--
	}
	return node.Val
}