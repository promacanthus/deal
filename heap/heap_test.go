package heap

import (
	"reflect"
	"testing"
)

var origin = []int{0, 33, 27, 21, 16, 13, 15, 19, 5, 6, 7, 8, 1, 2, 12}

func TestHeap_Sort(t *testing.T) {
	type fields struct {
		items []int
		los   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				items: InitHeap(origin, false).items,
				los:   false,
			},
			[]int{0, 1, 2, 5, 6, 7, 8, 12, 13, 15, 16, 19, 21, 27, 33},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Heap{
				items: tt.fields.items,
				los:   tt.fields.los,
			}
			if got := h.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
