package sort

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test bucket sort",
			args{list: []int{5, 4, 6, 2, 3, 1}},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BucketSort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMax(t *testing.T) {
	type args struct {
		list   []int
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test get max",
			args{
				list:   []int{1, 2, 3, 4, 5},
				length: 5,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMax(tt.args.list, tt.args.length); got != tt.want {
				t.Errorf("getMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
