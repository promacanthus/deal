package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{4, 5, 1, 6, 2, 3}
	type args struct {
		list []int
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test bubble sort",
			args: args{a},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.list)
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{4, 5, 1, 6, 2, 3, 4, 5, 1, 6, 2, 3, 4, 5, 1, 6, 2, 3})
	}
}
