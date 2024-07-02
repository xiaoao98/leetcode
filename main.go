package main

import (
	"fmt"
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
	// startGene := "AACCGGTT"
	// endGene := "AACCGGTA"
	// bank := []string{"AACCGGTA"}
	// fmt.Println(daily.MinMutation(startGene, endGene, bank))
	// a := [][]int{{1,2}, {3, 1}, {4,5}}
	// sort.Slice(a, func(i, j int)bool{
	// 	return a[i][1] < a[j][1]
	// })
	// fmt.Println(a)
	// charArr := []byte("asdafesfqa")
	// sort.Slice(charArr, func(i, j int)bool {
	// 	return charArr[i] < charArr[j]
	// })
	// fmt.Println(string(charArr))
	// s := []int{3,4,8,6}
	// s = append(s, 7)
	// sort.Ints(s)
	// fmt.Println(s)
	// rand.Seed(time.Now().Unix())
	// a := rand.Intn(30)
	// fmt.Println(a)
	// sort.Slice(s, func(i,j int)bool {
	// 	return s[i] > s[j]
	// })
	// fmt.Println(s)
	// c1 := 'c'
	// c2 := 'a'
	// fmt.Println(int(c1-c2))
	// if (c1 >= 'a' && c1 <= 'z') || (c1 >= 'A' && c1 <= 'Z') || (c1 >= '0' && c1 <= '9') {
	// 	fmt.Println(string(c1))
	// } 
	// s1 := "425"
	// val1, _ := strconv.Atoi(s1)
	// fmt.Println(val1 + 3)
	// fmt.Println(strconv.Itoa(val1)+"3")
	m1 := map[int]int{3: 4, 5:6, 7: 8}
	fmt.Println(len(m1))
	for key, val := range m1 {
		fmt.Println(key, val)
		delete(m1, key)
	}
	fmt.Println(len(m1))
}