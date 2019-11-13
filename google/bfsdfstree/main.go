package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func Insert(root **Node, value int) {
	if *root == nil {
		*root = &Node{Value: value}
		return
	}
	
	if value < (*root).Value {
		Insert(&(*root).Left, value)
	} else {
		Insert(&(*root).Right, value)
	}
}

func PrintDFS(root *Node) {
	if root == nil {
		return
	}
	PrintDFS(root.Left)
	fmt.Printf("-> %d ", root.Value)
	PrintDFS(root.Right)
}

type LLNode struct {
	Value *Node
	Next *LLNode
}

func PrintBFS(root *Node) {
	stack := &LLNode{Value:root}
	for stack != nil {
		var stack2 *LLNode
		for cursor := stack; cursor != nil; cursor = cursor.Next {
			node := cursor.Value
			fmt.Printf("--> %d ", node.Value)
			if node.Left != nil {
				stack2 = &LLNode{
					Value: node.Left,
					Next: stack2, 
				}
			}
			if node.Right != nil {
				stack2 = &LLNode{
					Value: node.Right,
					Next: stack2, 
				}
			}
		}
		fmt.Println()
		stack = stack2
	}
}

func main() {
	var root *Node

	Insert(&root, 50)
	Insert(&root, 25)
	Insert(&root, 75)
	Insert(&root, 12)
	Insert(&root, 37)
	Insert(&root, 62)
	Insert(&root, 77)
	Insert(&root, 55)
	Insert(&root, 5)
	Insert(&root, 13)
	Insert(&root, 36)
	
	// PrintDFS(root)
	PrintBFS(root)
}
