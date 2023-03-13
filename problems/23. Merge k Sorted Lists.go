package problems

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{}
	resB := res

	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			lists[i], lists[len(lists)-1] = lists[len(lists)-1], lists[i]
			lists = lists[:len(lists)-1]
			i--
		}
	}

	for len(lists) != 0 {
		min := math.MaxInt
		idx := 0
		for i := range lists {
			if lists[i].Val < min {
				min = lists[i].Val
				idx = i
			}
		}
		res.Next = lists[idx]
		res = res.Next
		lists[idx] = lists[idx].Next
		if lists[idx] == nil {
			lists[idx], lists[len(lists)-1] = lists[len(lists)-1], lists[idx]
			lists = lists[:len(lists)-1]
		}
	}
	return resB.Next
}
