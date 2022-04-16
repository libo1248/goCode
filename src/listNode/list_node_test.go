package main

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1, l2 := &ListNode{1, nil}, &ListNode{1, nil}
	tl1, tl2 := l1, l2
	for _, n := range []int{4, 5} {
		tl1.Next = &ListNode{n, nil}
		tl1 = tl1.Next
	}
	for _, n := range []int{2, 3} {
		tl2.Next = &ListNode{n, nil}
		tl2 = tl2.Next
	}
	l := addTwoNumbers(l1, l2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
