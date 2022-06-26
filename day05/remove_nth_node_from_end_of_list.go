/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 주어진 숫자대로 뒤에서부터 카운팅하여 노드 제거
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nth := head
	last := head

	for i := 0; i < n+1; i++ {
		if last == nil {
			return head.Next
		}
		last = last.Next
	}

	for last != nil {
		nth = nth.Next
		last = last.Next
	}

	nth.Next = nth.Next.Next

	return head
}