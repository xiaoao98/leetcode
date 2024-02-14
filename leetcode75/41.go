package leetcode75

import "math"

func maxLevelSum(root *TreeNode) int {
	maxLevel := math.MinInt
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	ret := 0
	for len(queue) > 0 {
		level += 1
		size := len(queue)
		levelNum := 0
		for i := 0; i < size; i++ {
			levelNum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		if levelNum > maxLevel {
			maxLevel = levelNum
			ret = level
		}
	}
	return ret
}