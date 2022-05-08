package sol

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	arr := []*ListNode{}
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	nLen := len(arr)
	if nLen <= 1 || k == 1 {
		return head
	}
	m := nLen / k
	for idx := 0; idx < m; idx++ {
		ReverseList(&arr, idx*k, (idx+1)*k-1, k)
	}
	head = arr[k-1]
	return head
}

func ReverseList(arr *[]*ListNode, start int, end int, step int) {
	temp := *arr
	for idx := end; idx >= start; idx-- {
		if idx-1 >= 0 {
			temp[idx].Next = temp[idx-1]
		}
	}
	if start+2*step-1 < len(temp) {
		temp[start].Next = temp[start+2*step-1]
	} else {
		if end+1 < len(temp) {
			temp[start].Next = temp[end+1]
		} else {
			temp[start].Next = nil
		}
	}
}
