package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func newNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func (n *ListNode) setNext(node *ListNode) *ListNode {
	n.Next = node
	return n
}

func main() {
	var l1, l2 *ListNode
	addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	start := head
	dec := 0
	sum := 0
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + dec
		dec = sum / 10
		sum = sum % 10
		if head == nil {
			head = &ListNode{Val: sum, Next: nil}
			start = head
		} else {
			head.Next = &ListNode{Val: sum, Next: nil}
			head = head.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	var tail *ListNode

	if l1 != nil {
		tail = l1
	} else {
		tail = l2
	}

	for tail != nil {
		sum = tail.Val + dec
		dec = sum / 10
		sum = sum % 10
		head.Next = &ListNode{Val: sum, Next: nil}
		head = head.Next
		tail = tail.Next
	}

	if dec == 1 {
		head.Next = &ListNode{Val: dec, Next: nil}
		head = head.Next
	}

	return start
}
