package substring_test

import (
	"testing"

	"github.com/lroolle/deepdivego/testing/substring"
)

func TestLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ABC", args{"abcabc"}, 3},
		{"中文", args{"我爱你你爱我"}, 3},
		{"Edge", args{""}, 0},
		{"Blank", args{" "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substring.LongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestsubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestSubstring2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ABC", args{"abcabc"}, 3},
		{"BBB", args{"bbb"}, 1},
		{"中文", args{"我爱你你爱我"}, 3},
		{"Edge", args{""}, 0},
		{"Blank", args{" "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substring.LongestSubstring2(tt.args.s); got != tt.want {
				t.Errorf("longestSubstring2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLongestSubstring(b *testing.B) {
	s := "黑化肥挥发发灰会挥发灰化肥挥发发黑会发挥"
	for i := 0; i < 10; i++ {
		s = s + s
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		substring.LongestSubstring(s)
	}
}

func BenchmarkLongestSubstring2(b *testing.B) {
	s := "黑化肥挥发发灰会挥发灰化肥挥发发黑会发挥"
	for i := 0; i < 10; i++ {
		s = s + s
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		substring.LongestSubstring2(s)
	}
}
