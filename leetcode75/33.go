package leetcode75

import "fmt"

func PairSum(head *ListNode) int {
	dummy := &ListNode{Val: 0, Next: head}
	slow, fast := dummy, dummy
	var last *ListNode
	for fast.Next != nil {
		fast = fast.Next.Next
		tmp := slow.Next
		slow.Next = last
		last = slow
		slow = tmp
	}
	right := slow.Next
	slow.Next = last
	ret := 0
	for right != nil {
		fmt.Println(right.Val, slow.Val)
		if right.Val+slow.Val > ret {
			ret = right.Val + slow.Val
		}
		right = right.Next
		slow = slow.Next
	}
	return ret
}