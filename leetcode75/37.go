package leetcode75

func pathSum(root *TreeNode, targetSum int) int {
	valueStored := make([]int, 0)
	return getPathSum(root, valueStored, targetSum)
}

func getPathSum(node *TreeNode, valueStored []int, targetSum int) int {
	if node == nil {
		return 0
	}
	sumNow := 0
	valueStored = append(valueStored, 0)
	for key, val := range valueStored {
		valueStored[key] += node.Val
		if val+node.Val == targetSum {
			sumNow += 1
		}
	}
	leftSum := getPathSum(node.Left, valueStored, targetSum)
	rightSum := getPathSum(node.Right, valueStored, targetSum)
	for key := range valueStored {
		valueStored[key] -= node.Val
	}
	return sumNow + leftSum + rightSum
}
