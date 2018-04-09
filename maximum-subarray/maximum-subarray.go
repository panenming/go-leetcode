package leetcode

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	temp := nums[0]
	max := temp
	i := 1
	for i < l {
		if temp < 0 {
			temp = nums[i]
		} else {
			temp += nums[i]
		}

		if max < temp {
			max = temp
		}
		i++
	}

	return max
}
