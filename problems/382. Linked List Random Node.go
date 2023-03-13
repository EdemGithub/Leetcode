package problems

import (
	"math/rand"
	"time"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Solution struct {
	sl []int
}

func Constructor(head *ListNode) Solution {
	rand.Seed(time.Now().UnixNano())
	s := Solution{}
	for head != nil {
		s.sl = append(s.sl, head.Val)
		head = head.Next
	}
	return s
}

func (this *Solution) GetRandom() int {
	return this.sl[rand.Intn(len(this.sl))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
