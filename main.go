package main

import (
	"fmt"
	"leetcode/daily"
	"sort"
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
	// s := "1101"
	// fmt.Println(daily.NumSteps(s))
	startGene := "AACCGGTT"
	endGene := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	fmt.Println(daily.MinMutation(startGene, endGene, bank))
	// a := [][]int{{1,2}, {3, 1}, {4,5}}
	// sort.Slice(a, func(i, j int)bool{
	// 	return a[i][1] < a[j][1]
	// })
	// fmt.Println(a)
	charArr := []byte("asdafesfqa")
	sort.Slice(charArr, func(i, j int)bool {
		return charArr[i] < charArr[j]
	})
	fmt.Println(string(charArr))
}