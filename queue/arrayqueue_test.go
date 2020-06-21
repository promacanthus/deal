package queue

import (
	"reflect"
	"testing"
)

func TestArrayQueue_Dequeue(t *testing.T) {
	type fields struct {
		items    []int
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
			a := &ArrayQueue{
				items:    tt.fields.items,
				length:   tt.fields.length,
				capacity: tt.fields.capacity,
			}
			got, got1 := a.Dequeue()
			if got != tt.want {
				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Dequeue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArrayQueue_Enqueue(t *testing.T) {
	type fields struct {
		items    []int
		length   uint
		capacity uint
	}
	type args struct {
		elem int
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
			a := &ArrayQueue{
				items:    tt.fields.items,
				length:   tt.fields.length,
				capacity: tt.fields.capacity,
			}
			if got := a.Enqueue(tt.args.elem); got != tt.want {
				t.Errorf("Enqueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewArrayQueue(t *testing.T) {
	type args struct {
		capacity uint
	}
	tests := []struct {
		name string
		args args
		want *ArrayQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayQueue(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
