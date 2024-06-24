package daily

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 type ListNode struct {
    Val int
    Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummyHead := &ListNode{Val: 0, Next: head}
	leftLast := dummyHead
	for i := 1; i < left; i ++ {
		leftLast = leftLast.Next
	}
	leftNode := leftLast.Next
	node := leftNode
	nextNode := node.Next
	for i:= left; i < right; i ++ {
		tmp := nextNode.Next
		nextNode.Next = node
		node = nextNode
		nextNode = tmp
	}
	leftLast.Next = node
	leftNode.Next = nextNode
	return dummyHead.Next
}