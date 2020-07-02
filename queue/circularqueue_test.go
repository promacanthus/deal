package queue

import (
	"reflect"
	"testing"
)

func TestCircularQueue_Dequeue(t *testing.T) {
	type fields struct {
		items    []int
		head     uint
		tail     uint
		capacity uint
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CircularQueue{
				items:    tt.fields.items,
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				capacity: tt.fields.capacity,
			}
			got, got1 := c.Dequeue()
			if got != tt.want {
				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Dequeue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCircularQueue_Enqueue(t *testing.T) {
	type fields struct {
		items    []int
		head     uint
		tail     uint
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
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CircularQueue{
				items:    tt.fields.items,
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				capacity: tt.fields.capacity,
			}
			if got := c.Enqueue(tt.args.elem); got != tt.want {
				t.Errorf("Enqueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCircularQueue(t *testing.T) {
	type args struct {
		capacity uint
	}
	tests := []struct {
		name string
		args args
		want *CircularQueue
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCircularQueue(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCircularQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
