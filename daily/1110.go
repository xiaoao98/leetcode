package daily

func delNodes(root *TreeNode, to_delete []int) []*TreeNode { //in recursion, you can use the node as return to delete or not
	ret := make([]*TreeNode, 0)
	toDeleteMap := make(map[int]bool, 0)
	for _, val := range to_delete {
		toDeleteMap[val] = true
	}
	var helper func(*TreeNode, bool)
	helper = func(node *TreeNode, isNewRoot bool) {
		if node == nil {
			return
		}
		if toDeleteMap[node.Val] {
			tmp1 := node.Left
			tmp2 := node.Right
			node.Left = nil
			node.Right = nil
			helper(tmp1, true)
			helper(tmp2, true)
		} else {
			if isNewRoot {
				ret = append(ret, node)
			}
			helper(node.Left, false)
			helper(node.Right, false)
		}
		if node.Left != nil && toDeleteMap[node.Left.Val] {
			node.Left = nil
		}
		if node.Right != nil && toDeleteMap[node.Right.Val] {
			node.Right = nil
		}
	}
	helper(root, true)
	return ret
}