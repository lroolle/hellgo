package substring

func LongestSubstring(s string) int {
	var start = 0
	var ret = 0
	var m = map[rune]int{}

	for i, ch := range []rune(s) {
		if lastI, ok := m[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > ret {
			ret = i - start + 1
		}
		m[ch] = i
	}
	return ret
}

func LongestSubstring2(s string) int {
	var start = 0
	var ret = 0
	// stores last occured pos + 1
	// 空间换时间，长串更省时间
	var lastOccur = make([]int, 0xffff)
	for i := range lastOccur {
		lastOccur[i] = -1
	}

	for i, ch := range []rune(s) {
		if lastI := lastOccur[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > ret {
			ret = i - start + 1
		}
		lastOccur[ch] = i
	}
	return ret
}
