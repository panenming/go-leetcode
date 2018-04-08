package leetcode

func searchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) && nums[i] <= target {
		if nums[i] == target {
			return i
		}

		i++
	}

	return i
}

func searchInsert2(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var media int
	// flag 用于标识最后一次循环是high下降还是low上升
	// 如果是high下降则直接返回media，如果是low上升则需要将media + 1，
	// 因为low上升代表nums[media] < target 这时需要进行media + 1
	var flag bool
	for low <= high {
		media = (low + high) / 2
		switch {
		case nums[media] < target:
			low = media + 1
			flag = true
		case nums[media] > target:
			high = media - 1
			flag = false
		default:
			return media
		}
	}
	if flag {
		return media + 1
	} else {
		return media
	}
}
