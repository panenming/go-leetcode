package leetcode

func singleNumber(nums []int) int {
	res := 0
	for _, q := range nums {
		res ^= q
	}

	return res
}
