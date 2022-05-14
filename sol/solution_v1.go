package sol

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroupV2(head *ListNode, k int) *ListNode {
	nLen := 0
	cur := head
	for cur != nil {
		nLen++
		cur = cur.Next
	}
	if nLen <= 1 || k == 1 {
		return head
	}
	m := nLen / k
	max := m * k
	cur = head
	var tail, preTail, prevHead *ListNode
	startCount := 0
	for startCount < max {
		cur, tail, prevHead, startCount = ReverseKNode(cur, startCount, k)
		if startCount == k {
			head = prevHead
		}
		if preTail != nil {
			preTail.Next = prevHead
		}
		preTail = tail
	}
	if cur != nil {
		tail.Next = cur
	}
	return head
}

func ReverseKNode(startNode *ListNode, startCount int, step int) (*ListNode, *ListNode, *ListNode, int) {
	var tail, prev *ListNode
	cur := startNode
	count := startCount
	max := startCount + step
	for count < max {
		if prev == nil {
			tail = cur
			prev = tail
			cur = cur.Next
			prev.Next = nil
		} else {
			temp := prev
			prev = cur
			cur = cur.Next
			prev.Next = temp
		}
		count++
	}
	return cur, tail, prev, count
}
