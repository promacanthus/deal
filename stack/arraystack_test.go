package stack

import (
	"reflect"
	"testing"
)

func TestArrayStack_Pop(t *testing.T) {
	type fields struct {
		data     []int
		length   uint
		capacity uint
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ArrayStack{
				items:    tt.fields.data,
				length:   tt.fields.length,
				capacity: tt.fields.capacity,
			}
			got, got1 := s.Pop()
			if got != tt.want {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArrayStack_Push(t *testing.T) {
	type fields struct {
		data     []int
		length   uint
		capacity uint
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ArrayStack{
				items:    tt.fields.data,
				length:   tt.fields.length,
				capacity: tt.fields.capacity,
			}
			if got := s.Push(tt.args.n); got != tt.want {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStack(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want *ArrayStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
