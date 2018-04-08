package leetcode

func searchRange(nums []int, target int) []int {
	// 查看target是否在nums中
	index := search(nums, target)

	if index == -1 {
		return []int{-1, -1}
	}

	fisrt := index
	for {
		f := search(nums[:fisrt], target)
		if f == -1 {
			break
		}
		fisrt = f
	}

	last := index
	for {
		l := search(nums[last+1:], target)
		if l == -1 {
			break
		}
		last = last + l + 1
	}

	return []int{fisrt, last}
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var media int

	for low <= high {
		media = (low + high) / 2
		switch {
		case nums[media] < target:
			low = media + 1
		case nums[media] > target:
			high = media - 1
		default:
			return media
		}

	}
	return -1
}
