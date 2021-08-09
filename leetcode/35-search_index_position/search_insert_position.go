package searchindexposition

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if target <= nums[0] {
			return 0
		}
		return 1
	}

	index := 0
	for i := 0; i < len(nums); i++ {
		index = i
		if nums[i] >= target {
			break
		}
	}

	if index == 0 {
		return 0
	}
	if nums[index] == target || (nums[index] > target && nums[index-1] < target) {
		return index
	}
	return index + 1
}
