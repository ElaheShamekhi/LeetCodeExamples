package LengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	var result int
	charIndexMap := make(map[uint8]int)
	start := 0
	for end := 0; end < n; end++ {
		if duplicateIndex, isDuplicate := charIndexMap[s[end]]; isDuplicate {
			result = max(result, end-start)
			start = duplicateIndex + 1
		}
		charIndexMap[s[end]] = end
	}
	result = max(result, n-start)

	return result
}
