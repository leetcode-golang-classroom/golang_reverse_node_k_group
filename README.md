# golang_reverse_node_k_group

Given the `head` of a linked list, reverse the nodes of the list `k` at a time, and return *the modified list*.

`k` is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of `k` then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

## Examples

Given the `head` of a linked list, reverse the nodes of the list `k` at a time, and return *the modified list*.

`k` is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of `k` then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

**Example 1:**

![https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg)

```
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg)

```
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]

```

**Constraints:**

- The number of nodes in the list is `n`.
- `1 <= k <= n <= 5000`
- `0 <= Node.val <= 1000`

**Follow-up:** Can you solve the problem in `O(1)` extra memory space?

## 解析

給定一個單向鏈結串列 list 還有一個整數 k

要求寫一個演算法把每 k 個一組的結點做反序

作法一：

 先用一個指標陣列把每個走訪過的結點都存起來

然後算出 nLen/k

每次取出 k個做反向

這樣空間複雜度跟時間複雜度都是 O(n)

作法二：

 如果依照順序，每次把新的結點放在前一個結點的的前面

這樣結果也會是反序

要這樣做必須每次紀錄上一次走訪到的結點

作法是先算出 nLen/k

也就是最多要處理 nLen/k * k個結點

每次針對 k個結點做反向

## 程式碼- 作法一

```go
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
```

## 程式碼-作法2

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
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
		if tail == nil {
			tail = cur
		}
		if prev == nil {
			prev = cur
			cur = cur.Next
		} else {
			temp := prev
			prev = cur
			cur = cur.Next
			prev.Next = temp
		}
		count++
	}
	tail.Next = nil
	return cur, tail, prev, count
}
func Transverse(head *ListNode) {
	cur := head
	for cur != nil {
		log.Printf("%v->", cur.Val)
		cur = cur.Next
	}
	log.Println()
}
```

## 困難點

1. 因為是單向鏈結，所以要反序必須要紀錄每一個前一個結點
2. 要重複 nLen/k 次反序
3. 每次要紀錄上一次的tail, head 必須知道如何接續

## Solve Point

- [x]  Understand what problem would like to solve
- [x]  Analysis Complexity