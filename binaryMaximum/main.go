// https://leetcode.com/problems/binary-tree-maximum-path-sum/
package main

func maxPathWithRoot(root *TreeNode) (bestWithRoot, best int) {
    var left, leftBest, right, rightBest int
    
    bestWithRoot = root.Val
    
    if (root.Left != nil) {
        left, leftBest = maxPathWithRoot(root.Left)
    }
    if root.Val + left > bestWithRoot {
        bestWithRoot = root.Val + left
    }
    
    if (root.Right != nil) {
        right, rightBest = maxPathWithRoot(root.Right)
    }
    if root.Val + right > bestWithRoot {
        bestWithRoot = root.Val + right
    }
    
    best += root.Val
    if left > 0 {
        best += left
    }
    if right > 0 {
        best += right
    }
    if root.Left != nil && leftBest > best {
        best = leftBest
    }
    if root.Right != nil && rightBest > best {
        best = rightBest
    }
    
    return bestWithRoot, best
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    _, best := maxPathWithRoot(root)
    return best
}

func main() {
    fmt.Println("hello mate")
}
