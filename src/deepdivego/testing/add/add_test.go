package add_test

import (
	"testing"

	. "github.com/lroolle/deepdivego/testing/add"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Who the fuck told you 1+1=2
		{"1 + 1", args{1, 1}, 3},
		{"1 + 2", args{1, 2}, 3},
		{"1 + 2", args{1, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
			// NOTE: The dot notation do not import non-exported ~add~
			// if got := Add(tt.args.a, tt.args.b); got != tt.want {
			// 	t.Errorf("Add() = %v, want %v", got, tt.want)
			// }
		})
	}
}
