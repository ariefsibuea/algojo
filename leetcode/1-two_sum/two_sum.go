package twosum

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := range nums {
		complement := target - nums[i]
		if _, ok := m[complement]; ok {
			return []int{m[complement], i}
		}

		m[nums[i]] = i
	}
	return nil
}
