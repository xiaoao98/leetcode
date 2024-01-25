package leetcode75

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ret := make([]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		ret = append(ret, queue[size-1].Val)
		for i := 0; i < size; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
	}
	return ret
}