package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test insertion sort",
			args{[]int{4, 5, 1, 6, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.list)
			fmt.Println(tt.args.list)
		})
	}
}
