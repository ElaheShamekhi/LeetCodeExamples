package LongestAscendingSubarray

func longestAscendingSubarray(arr []int) int {
	current, max := 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			current++
		} else {
			if current > max {
				max = current
			}
			current = 1
		}
	}
	if current > max {
		max = current
	}
	return max
}
