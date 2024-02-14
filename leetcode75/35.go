package leetcode75

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaf1 []int
	var leaf2 []int
	leaf1 = findAllLeaves(root1)
	leaf2 = findAllLeaves(root2)
	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}

func findAllLeaves(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	return append(findAllLeaves(root.Left), findAllLeaves(root.Right)...)
}