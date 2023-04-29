package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode = nil
	for head != nil {
		result = &ListNode{head.Val, result}
		head = head.Next
	}
	return result
}
