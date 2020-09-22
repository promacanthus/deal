package linkedlists

import "testing"

func Test_hasCycle1(t *testing.T) {
	a := &ListNode{
		Val:  1,
		Next: nil,
	}
	a.Next = a
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{head: a},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle1(tt.args.head); got != tt.want {
				t.Errorf("hasCycle1() = %v, want %v", got, tt.want)
			}
		})
	}
}
