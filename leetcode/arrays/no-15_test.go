package arrays

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{[]int{-1, 0, 1, 2, -1, -1}},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSumDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
