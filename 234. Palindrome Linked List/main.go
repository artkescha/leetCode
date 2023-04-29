package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	reverseHead := reverseList(head)
	for head != nil && reverseHead != nil {
		if head.Val != reverseHead.Val {
			return false
		}
		head = head.Next
		reverseHead = reverseHead.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode = nil
	for head != nil {
		result = &ListNode{head.Val, result}
		head = head.Next
	}
	return result
}
