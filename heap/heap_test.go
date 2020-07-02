package heap

import (
	"fmt"
	"testing"
)

var origin = []int{0, 33, 27, 21, 16, 13, 15, 9, 5, 6, 7, 8, 1, 2}

func TestHeap_Delete(t *testing.T) {
	type fields struct {
		items []int
		count uint
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				items: origin,
				count: 13,
			},
			33,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				items: tt.fields.items,
				count: tt.fields.count,
			}
			got, got1 := h.Delete()
			if got != tt.want {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Delete() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHeap_Get(t *testing.T) {
	type fields struct {
		items []int
		count uint
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				items: origin,
				count: 13,
			},
			33,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				items: tt.fields.items,
				count: tt.fields.count,
			}
			got, got1 := h.Get()
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHeap_Insert(t *testing.T) {
	type fields struct {
		items []int
		count uint
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				items: origin,
				count: 13,
			},
			args{v: 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				items: tt.fields.items,
				count: tt.fields.count,
			}
			h.Insert(tt.args.v)
			fmt.Println(h.items)
		})
	}
}
