package sol

import (
	"reflect"
	"testing"
)

func CreateList(nums *[]int) *ListNode {
	var head, cur *ListNode
	arr := *nums
	for idx, val := range arr {
		if idx == 0 {
			head = &ListNode{Val: val}
			cur = head
		} else {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
	}
	return head
}
func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [1,2,3,4,5], k = 2",
			args: args{head: CreateList(&[]int{1, 2, 3, 4, 5}), k: 2},
			want: CreateList(&[]int{2, 1, 4, 3, 5}),
		},
		{
			name: "head = [1,2,3,4,5], k = 2",
			args: args{head: CreateList(&[]int{1, 2, 3, 4, 5}), k: 3},
			want: CreateList(&[]int{3, 2, 1, 4, 5}),
		},
		{
			name: "head = [1,2], k = 2",
			args: args{head: CreateList(&[]int{1, 2}), k: 2},
			want: CreateList(&[]int{2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
