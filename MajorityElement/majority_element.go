package MajorityElement

/*
	Given an array nums of size n, return the majority element.
	The majority element is the element that appears more than ⌊n / 2⌋ times.
	You may assume that the majority element always exists in the array.
*/

func MajorityElement(nums []int) int {
	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num] += 1
		if numsMap[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}
