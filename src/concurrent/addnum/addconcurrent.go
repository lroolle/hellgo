package main

import (
	"runtime"
	"sync"
)

func AddConcurrent(nums []int) int {
	var ret int
	gn := runtime.NumCPU()
	if len(nums) <= gn {
		return ret
	}
	// divide nums into gn(numbers of goroutines) groups
	// one goroutine in charge of count each part
	groupSize := len(nums) / gn

	var numsTmp []int
	var wg sync.WaitGroup
	wg.Add(gn)

	for i := 0; i < gn; i++ {
		go func(i int) {
			start := i * groupSize
			if i == len(nums)-1 {
				end := len(nums)
			} else {
				end := start + i
			}

		}(i)
		wg.Done()
	}

	wg.Wait()
	return ret
}
