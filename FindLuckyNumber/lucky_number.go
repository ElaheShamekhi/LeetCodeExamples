package FindLuckyNumber

/*
	Given an array of integers arr, a lucky integer is an integer that has a frequency in the array equal to its value.
	Return the largest lucky integer in the array. If there is no lucky integer return -1.
*/

func findLucky(arr []int) int {
	frequency := make(map[int]int)
	for _, num := range arr {
		frequency[num]++
	}
	maxLucky := -1
	for num, freq := range frequency {
		if num == freq && num > maxLucky {
			maxLucky = num
		}
	}
	return maxLucky
}
