package leetcode

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	sort.Ints(nums)
	max, temp := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			temp++
		} else if nums[i] != nums[i-1] {
			// 0,1,1,2 可以继续
			temp = 1
		}

		if max < temp {
			max = temp
		}
	}

	return max
}
