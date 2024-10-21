package main

import (
	_ "net/http/pprof"
)

func main() {

}

//给定两个用链表表示的整数，每个节点包含一个数位。
//这些数位是反向存放的，也就是个位排在链表首部。
//编写函数对这两个整数求和，并用链表形式返回结果。
//示例：
//输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912

type Node struct {
	Val  int
	Next *Node
}

func listNodeAdd(nodeA, nodeB *Node) *Node {
	if nodeA == nil {
		return nodeB
	}
	if nodeB == nil {
		return nodeA
	}
	overflow := 0
	val := 0
	head := new(Node)
	curr := head
	for nodeA != nil && nodeB != nil {
		val = (nodeA.Val + nodeB.Val + overflow) % 10
		overflow = (nodeA.Val + nodeB.Val) / 10
		node := new(Node)
		node.Val = val
		node.Next = nil
		curr.Next = node
		curr = node
		nodeA = nodeA.Next
		nodeB = node.Next
	}

	if nodeA != nil {
		curr.Next = nodeA
	}

	if nodeB != nil {
		curr.Next = nodeB
	}
	return head.Next
}
