package main

import "testing"

func TestSums(t *testing.T) {
	t.Run("Arrays", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Sums(nums)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, nums)
		}

	})

	t.Run("Slices", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Sums(nums)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, nums)
		}

	})
}

func TestSumAll(t *testing.T) {
	t.Run("SumAll", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2})
		want := []int{15, 3}
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
