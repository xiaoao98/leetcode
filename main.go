package main

import (
	"fmt"
	"leetcode/leetcode75"
)

//
func main() {
	// s := "3[a]2[bc]"
	// fmt.Print(leetcode75.DecodeString(s))
	a := &leetcode75.ListNode{Val: 1, Next: nil}
	b := &leetcode75.ListNode{Val: 2, Next: nil}
	a.Next = b
	output := leetcode75.ReverseList(a)
	fmt.Print(output.Val, output.Next.Val, output.Next)
}