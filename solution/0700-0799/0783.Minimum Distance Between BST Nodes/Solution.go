/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res int
var preNode *TreeNode
func minDiffInBST(root *TreeNode) int {
    res = int(^uint(0) >> 1)
    preNode = nil
    helper(root)
    return res
}

func helper(root *TreeNode)  {
    if root == nil {
        return
    }
    helper(root.Left)
    if preNode != nil {
        res = getMinInt(res, root.Val - preNode.Val)
    }
    preNode = root
    helper(root.Right)
}

func getMinInt(a,b int) int {
    if a < b {
        return a
    }
    return b
}