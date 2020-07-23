package array

import (
	"fmt"
	"testing"
)

func TestArray_Delete(t *testing.T) {
	type fields struct {
		data   []int
		length uint
	}
	type args struct {
		index uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			"test_delete",
			fields{
				data:   []int{1, 2, 3},
				length: 3,
			},
			args{index: uint(2)},
			3,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			got, err := a.Delete(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_Find(t *testing.T) {
	type fields struct {
		data   []int
		length uint
	}
	type args struct {
		index uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			"test_find",
			fields{
				data:   []int{1, 2, 3},
				length: 3,
			},
			args{index: uint(1)},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			got, err := a.Find(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_Insert(t *testing.T) {
	type fields struct {
		data   []int
		length uint
	}
	type args struct {
		index uint
		v     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test_insert",
			fields{
				data:   []int{1, 2, 3},
				length: 3,
			},
			args{
				index: uint(3),
				v:     4,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Array{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if err := a.Insert(tt.args.index, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArray_InsertToTail(t *testing.T) {
	type fields struct {
		data   []int
		length uint
	}
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test_inserttotail",
			fields{
				data:   make([]int, 3, 4),
				length: 0,
			},
			args{v: 4},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if err := a.InsertToTail(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("InsertToTail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArray_Len(t *testing.T) {
	type fields struct {
		data   []int
		length uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			"test_len",
			fields{
				data:   []int{1, 2, 3},
				length: 3,
			},
			uint(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if got := a.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_Print(t *testing.T) {
	type fields struct {
		data   []int
		length uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"test_print",
			fields{
				data:   []int{1, 2, 3},
				length: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			a.String()
		})
	}
}

func TestArray_isIndexOutOfRange(t *testing.T) {
	type fields struct {
		data   []int
		length uint
	}
	type args struct {
		index uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"test_isindexoutofrange",
			fields{
				data:   []int{1, 2, 3},
				length: 3,
			},
			args{index: uint(0)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if got := a.isIndexOutOfRange(tt.args.index); got != tt.want {
				t.Errorf("isIndexOutOfRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_expansion(t *testing.T) {
	type fields struct {
		data   []int
		length uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				data:   make([]int, 1),
				length: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			a.expand()
			fmt.Println(cap(a.data))
		})
	}
}
