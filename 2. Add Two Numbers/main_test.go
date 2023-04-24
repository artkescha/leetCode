package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	// l1 = 2->4->3
	l1 := newNode(2)
	l1.setNext(newNode(4).setNext(newNode(3)))
	// l2 = 5->6->4
	l2 := newNode(5)
	l2.setNext(newNode(6).setNext(newNode(4)))

	// l1+l2 = 7->0->8
	want1 := newNode(7)
	want1.setNext(newNode(0).setNext(newNode(8)))

	// l3 = 0
	l3 := newNode(0)
	// l4 = 0
	l4 := newNode(0)
	// l3+l4 = 0
	want2 := newNode(0)

	// l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	l5 := newNode(9)
	l5.setNext(newNode(9).setNext(newNode(9).setNext(newNode(9).setNext(newNode(9).setNext(newNode(9).setNext(newNode(9)))))))

	l6 := newNode(9)
	l6.setNext(newNode(9).setNext(newNode(9).setNext(newNode(9))))

	// l5+l6 = 8->9->9->9->0->0->0->1
	want3 := newNode(8)
	want3.setNext(newNode(9).setNext(newNode(9).setNext(newNode(9).setNext(newNode(0).
		setNext(newNode(0).setNext(newNode(0).setNext(newNode(1))))))))
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "happy path",
			// l1 = 2->4->3
			l1: l1,
			// l2 = 5->6->4
			l2:   l2,
			want: want1,
		},
		{
			name: "zero value",
			// l1 = 0
			l1: l3,
			// l2 = 0
			l2:   l4,
			want: want2,
		},
		{
			name: "lists of different lengths",
			// l1 = 0
			l1: l5,
			// l2 = 0
			l2:   l6,
			want: want3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := addTwoNumbers(tt.l1, tt.l2)
			if got := addTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", r, tt.want)
			}
		})
	}
}
