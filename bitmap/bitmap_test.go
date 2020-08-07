package bitmap

import (
	"reflect"
	"testing"
)

func TestNewBitMap(t *testing.T) {
	type args struct {
		bits int
	}
	tests := []struct {
		name string
		args args
		want *BitMap
	}{
		// TODO: Add test cases.
		{
			"1",
			args{10},
			&BitMap{
				bytes: make([]byte, 10/8+1),
				nbits: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBitMap(tt.args.bits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBitMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
