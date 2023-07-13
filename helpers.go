package main

func LengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	startMax := 0
	endMax := 0
	tmpStart := 0
	tmpIndex := 0
	for i := 1; i < len(s); i++ {
		tmpIndex = -1
		for j := tmpStart; j < i; j++ {
			if s[j] == s[i] {
				tmpIndex = j
				break
			}
		}
		if tmpIndex >= 0 {
			if (i - tmpStart - tmpIndex) > (endMax - startMax) {
				startMax = tmpIndex
				endMax = i
			}
			tmpStart = tmpIndex + 1
		} else {
			if (i - tmpStart + 1) > (endMax - startMax) {
				startMax = tmpStart
				endMax = i + 1
			}
		}
	}
	return endMax - startMax
}
