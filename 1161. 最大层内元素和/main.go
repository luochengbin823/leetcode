package main

import "leetcode/utils"

func maxLevelSum(root *utils.TreeNode) int {
	/*
		思路：
		用层序遍历，遍历每一层。将每一层的元素求和，比较得出最大值的层数即可。
		由于我们需要知道层的元素和，以及层数，因此我们用hashmap来进行存储
	*/
	//第1层为root的值
	record := make(map[int]int)
	record[root.Val] = 1
	//层序遍历套路
	//定义队列辅助遍历
	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	maxvalue := -1000000
	level := 0
	anslevel := 0
	for len(queue) != 0 {
		//先看当前队列长度，然后以此控制遍历
		levelsize := len(queue)
		levelcount := 0

		for i := 0; i < levelsize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelcount += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
		if levelcount > maxvalue {
			maxvalue = levelcount
			anslevel = level
		}

	}
	return anslevel
}
