package main

import "fmt"

/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：
输入：lists = []
输出：[]

示例 3：
输入：lists = [[]]
输出：[]

提示：
k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
*/

func main() {
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
	list1 := mergeKLists([]*ListNode{l1, l2})
	for list1 != nil {
		fmt.Println(list1.Val)
		list1 = list1.Next
	}

	list2 := mergeKLists1([]*ListNode{l1, l2})
	for list2 != nil {
		fmt.Println(list2.Val)
		list2 = list2.Next
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var l, tl *ListNode
	var minIndex int
	var minNode *ListNode
	var endCount int
	allEnd := false
	count := len(lists)
	for !allEnd {
		endCount = 0
		minNode = nil
		minIndex = 0

		for i, node := range lists {
			if node == nil {
				endCount++
				continue
			}

			if minNode == nil {
				minIndex = i
				minNode = node
				continue
			}

			if minNode.Val > node.Val {
				minIndex = i
				minNode = node
			}
		}

		if endCount == count {
			allEnd = true
			break
		}

		lists[minIndex] = lists[minIndex].Next
		if l == nil {
			l = minNode
			tl = l
		} else {
			tl.Next = minNode
			tl = tl.Next
		}
	}

	return l
}

func mergeKLists1(lists []*ListNode) *ListNode {
	count := len(lists)

	if count == 0 {
		return nil
	} else {
		return marge(lists, 0, count-1)
	}
}

func marge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	} else {
		mid := left + (right-left)/2
		l1 := marge(lists, left, mid)
		l2 := marge(lists, mid+1, right)
		return margeTwoLists(l1, l2)
	}
}

func margeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = margeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = margeTwoLists(list1, list2.Next)
		return list2
	}
}
