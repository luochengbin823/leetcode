package _50_删除二叉搜索树中的节点

import "leetcode/utils"

//删除二叉搜索树中的节点
func deleteNode(root *utils.TreeNode, key int) *utils.TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		return deleteNode(root.Left, key)
	case root.Val < key:
		return deleteNode(root.Right, key)
		//左子树为空或右子树为空的情况
	case root.Left == nil || root.Right == nil:
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	//左右子树都为空的情况
	default:
		//左右子树都存在的情况。要找比root大的最小节点，必然在右子树（因为此时肯定是左右子树都存在），那么递归的找右子树的最左节点
		sucessor := root.Right
		for sucessor.Left != nil {
			sucessor = sucessor.Left
		}
		sucessor.Right = deleteNode(root.Right, sucessor.Val)
		sucessor.Left = root.Left
		return sucessor
	}
	return root
}
