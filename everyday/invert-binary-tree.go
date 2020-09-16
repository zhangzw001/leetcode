package main

import "fmt"

/*
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
备注:
这个问题是受到 Max Howell 的 原问题 启发的 ：

谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
通过次数123,209提交次数160,272

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/invert-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
		root.Left,root.Right = root.Right,root.Left
		invertTree(root.Left)
		invertTree(root.Right)

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

 func (root *TreeNode) GetLeafNode(t *TreeNode) {
 	if t != nil {
 		if t.Left == nil && t.Right == nil {
 			fmt.Printf("%d ",t.Val)
		}
		root.GetLeafNode(t.Left)
 		root.GetLeafNode(t.Right)
	}
 }
func main() {
	tree := &TreeNode{1,&TreeNode{11,nil,nil},&TreeNode{12,nil,nil}}
	invertTree(tree)
	tree.GetLeafNode(tree)
}