package goBinaryTree

/*值为-1对应空结点*/

//生成树结点
func genTreeNode(val int) *TreeNode {
	if val == -1 {
		return nil
	}
	node := new(TreeNode)
	node.Val = val
	return node
}

//层序构造二叉树
func BuildLevelOrder(vals []int) *TreeNode {
	num := len(vals)
	if num == 0 {
		return nil
	}
	//构造数组中所有元素对应的结点
	nodes := make([]*TreeNode, 0)
	for i := 0; i < num; i++ {
		nodes = append(nodes, genTreeNode(vals[i]))
	}
	//将所有结点构造成二叉树
	tmp := 0
	for i := 0; i < num; i++ {
		if nodes[i] == nil {
			tmp += 2
			continue
		}
		if 2*i+1-tmp >= num {
			break
		}
		nodes[i].Left = nodes[2*i+1-tmp]
		if 2*i+2-tmp >= num {
			break
		}
		nodes[i].Right = nodes[2*i+2-tmp]
	}
	return nodes[0]
}