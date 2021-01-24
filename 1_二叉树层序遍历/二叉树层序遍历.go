// 二叉树层序遍历
// 2种方式：非递归方式  递归方式

package main

import(
	"fmt"
)

func main(){
	//根节点
	var root *TreeNode
	// 层序遍历
	res:= levelOrder(root)
	fmt.Println(res)
}

// TreeNode  树节点
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}


//*************非递归方式（广度优先）************
func levelOrder(root *TreeNode) [][]int {
	res:=make([][]int,0)
	if root==nil{
		return res
	}
	q:=[]*TreeNode{root}
	//层级
	for i:=0;0<len(q);i++{
		res=append(res,[]int{})
		p:=[]*TreeNode{}
		//循环每层元素
		for j:=0;j<len(q);j++{
			//初始化每层
			res[i]=append(res[i],q[j].Val)
			if q[j].Left!=nil{
				p=append(p, q[j].Left)
			}
			if q[j].Right!=nil{
				p=append(p,q[j].Right)
			}
		}
	q=p
		
	}
	return res
}

//**********************


// ************递归方式*************
// levelOrder  层序遍历
// func levelOrder(root *TreeNode) [][]int {
//     res=make([][]int,0)
// 	bfs(root,0)
// 	return res
// }
// var res [][]int

// func bfs(node *TreeNode,level int){
// 	if node==nil{
// 		return 
// 	}

// 	if len(res)==level{
// 		res=append(res,[]int{})
// 	}
// 	res[level]=append(res[level],node.Val)
// 	bfs(node.Left,level+1)
// 	bfs(node.Right,level + 1)
// }
//******************************************


