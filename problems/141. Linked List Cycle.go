package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		fast, slow = fast.Next.Next, slow.Next
	}
	return false
}
