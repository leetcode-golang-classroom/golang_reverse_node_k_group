package sol

import (
	"reflect"
	"testing"
)

func Test_reverseKGroupV2(t *testing.T) {
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
			if got := reverseKGroupV2(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroupV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
