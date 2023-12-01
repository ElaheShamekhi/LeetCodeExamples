package TowSum

func TwoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for idx, num := range nums {
		if nextIdx, seen := numsMap[target-num]; seen {
			return []int{nextIdx, idx}
		}
		numsMap[num] = idx
	}
	return []int{}
}
