package removeduplicatefromsortedarray

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	index, lastNum := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == lastNum {
			continue
		}

		lastNum = nums[i]
		nums[index] = lastNum
		index++
	}

	return index
}
