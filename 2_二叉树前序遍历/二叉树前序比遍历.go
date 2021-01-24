package main

import(
	"fmt"
)

// TreeNode 树节点
type TreeNode struct{
	Val int
	Left *TreeNode  
	Right *TreeNode
}

func main(){
	 r:= preorderTrvaersal(nil)
	 fmt.Println(r)

	 var r1 []int
	 fmt.Println(r1)
}

//***************  前序遍历  递归算法
// var res []int 
// func preOrderTrvaersal(root *TreeNode) []int{
// 	res=make([]int,0)
// 	dfs(root)
// 	return res 
// }

// func dfs(node *TreeNode){

// 	if node==nil{
// 		return
// 	}

// 	res=append(res,node.Val)
// 	dfs(node.Left)
// 	dfs(node.Right)
// }

//**********************************

func preorderTrvaersal(root *TreeNode)(res []int){

	if root==nil{
		return res
	}

	res=append(res,root.Val)

	res=append(res,preorderTrvaersal(root.Left)...)
	res=append(res,preorderTrvaersal(root.Right)...)
	return 
}


