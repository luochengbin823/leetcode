package _99__二叉树的右视图

import (
	"leetcode/utils"
)

func rightSideView(root *utils.TreeNode) []int {
	//层次遍历。每当队伍长度为1的时候，那个元素就是需要的元素
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	//当队列为0，暂存节点也为空时遍历结束
	for len(queue) != 0 {
		var Nodeval int
		//遍历逻辑为当队列为空时，元素入队。然后元素出队，元素子节点入队（如果队伍此时不为空，则不用入队）
		levelsize := len(queue)
		// 用levelsize来取除了最后一个元素外的所有元素。以控制层数
		for i := 0; i < levelsize; i++ {
			node := queue[0]
			queue = queue[1:]
			Nodeval = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, Nodeval)
	}
	return ans
}
