package removeelement

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}

	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}

		nums[index] = nums[i]
		index++
	}

	return index
}
