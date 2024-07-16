package daily

import "strings"

func getDirections(root *TreeNode, startValue int, destValue int) string {
	var startStr string
	var destStr string
	var dirToNode func(*TreeNode, []byte)
	dirToNode = func(node *TreeNode, dirNow []byte) {
		if node == nil {
			return
		}
		if node.Val == startValue {
			startStr = string(dirNow)
		} else if node.Val == destValue {
			destStr = string(dirNow)
		}
		if startStr != "" && destStr != "" {
			return
		}
		dirNow = append(dirNow, 'L')
		dirToNode(node.Left, dirNow)
		dirNow = dirNow[:len(dirNow)-1]
		dirNow = append(dirNow, 'R')
		dirToNode(node.Right, dirNow)
		dirNow = dirNow[:len(dirNow)-1]
	}

	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val == startValue || node.Val == destValue {
			return node
		}
		leftLCA := dfs(node.Left)
		rightLCA := dfs(node.Right)
		if leftLCA == nil {
			return rightLCA
		} else if rightLCA == nil {
			return leftLCA
		} else {
			return node
		}
	}
	LCA := dfs(root)
	
	dirToNode(LCA, []byte{})
	return strings.Repeat("U", len(startStr)) + destStr
}
