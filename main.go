package main

import (
	"fmt"
	"leetcode/daily"
)

//
func main() {
	// s := "3[a]2[bc]"
	// fmt.Print(leetcode75.DecodeString(s))
	// a := &leetcode75.ListNode{Val: 1, Next: nil}
	// b := &leetcode75.ListNode{Val: 2, Next: nil}
	// a.Next = b
	// output := leetcode75.ReverseList(a)
	// fmt.Print(output.Val, output.Next.Val, output.Next)
	// a := &leetcode75.ListNode{Val: 4, Next: nil}
	// b := &leetcode75.ListNode{Val: 2, Next: nil}
	// c := &leetcode75.ListNode{Val: 2, Next: nil}
	// d := &leetcode75.ListNode{Val: 3, Next: nil}
	// a.Next = b
	// b.Next = c
	// c.Next = d
	// fmt.Print(leetcode75.PairSum(a))
	// input := "etddgeia"
	// output := daily.TransferString(input)
	// fmt.Println(output)
	// s := "abcd"
	// t := "bcdf"
	// i := 0
	//cost := int(math.Abs(float64(s[i] - t[i])))
	// fmt.Print(float64(int(s[i]) - int(t[i])))
	//43 47 52 56 66 72 again 
	s := "1101"
	fmt.Println(daily.NumSteps(s))
}