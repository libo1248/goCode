package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	tl := l
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			tl.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tl.Val += l2.Val
			l2 = l2.Next
		}
		tl.Val += carry

		if tl.Val >= 10 {
			tl.Val -= 10
			carry = 1
		} else {
			carry = 0
		}

		if l1 != nil || l2 != nil || carry != 0 {
			tl.Next = &ListNode{}
			tl = tl.Next
		}
	}

	return l
}
