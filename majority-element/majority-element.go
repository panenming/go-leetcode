package leetcode

func majorityElement(nums []int) int {
	x, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		switch {
		case x == nums[i]:
			count++
		case count > 0:
			count--
		default:
			x = nums[i]
			count = 1

		}
	}
	return x
}
