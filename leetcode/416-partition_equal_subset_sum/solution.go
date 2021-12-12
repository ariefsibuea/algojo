package partitionequalsubsetsum

// Problem source: https://leetcode.com/problems/partition-equal-subset-sum/submissions/
// Solve this problem using Dynamic Programming

func CanPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	// bottom up manner to find subset
	part := make([]bool, (sum/2)+1)
	for i := 0; i < len(nums); i++ {
		for j := sum / 2; j >= nums[i]; j-- {
			if part[j-nums[i]] || j == nums[i] {
				part[j] = true
			}
		}
	}

	return part[sum/2]
}
