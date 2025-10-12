package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/cooperstandard/leetcode/datastructures"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// csv style explitly marking

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil,"
	}
	return dfs(*root)
}

func dfs(root TreeNode) string {
	output := fmt.Sprintf("%d,", root.Val)

	if root.Left == nil {
		output += "nil,"
	} else {
		output += dfs(*root.Left)
	}

	if root.Right == nil {
		output += "nil,"
	} else {
		output += dfs(*root.Right)
	}

	return output
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := datastructures.Stack[string]{Items: strings.Split(data, ",")}
	slices.Reverse(nodes.Items)
	fmt.Println(nodes.Items)

	if nodes.IsEmpty() || nodes.Peek() == "nil" {
		return nil
	}

	val, _ := strconv.Atoi(nodes.Pop())
	root := &TreeNode{Val: val}

	fill(root, &nodes)
	return root
}

func fill(root *TreeNode, nodes *datastructures.Stack[string]) {
	if nodes.IsEmpty() {
		return
	}

	if nodes.Peek() == "nil" || len(nodes.Peek()) == 0 {
		nodes.Pop()
	} else {

		val, _ := strconv.Atoi(nodes.Pop())
		fmt.Println(nodes.Items)

		root.Left = &TreeNode{Val: val}

		fill(root.Left, nodes)
	}

	if nodes.IsEmpty() {
		return
	}

	if nodes.Peek() == "nil" || len(nodes.Peek()) == 0 {
		nodes.Pop()
	} else {

		val, _ := strconv.Atoi(nodes.Pop())
		fmt.Println(val)

		root.Right = &TreeNode{Val: val}

		fill(root.Right, nodes)
	}
}
