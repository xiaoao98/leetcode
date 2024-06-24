package daily

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		toAdd := make([]int, 0)
		toAdd = append(toAdd, level...)
		ret = append(ret, toAdd)
		queue = queue[size:]
	}
	return ret
}