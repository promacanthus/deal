package linkedlists

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			"1",
			args{lists: []*ListNode{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
			}},
			&ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists2(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
