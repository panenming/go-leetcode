package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, b := range nums {
		if j, ok := m[target-b]; ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	return nil
}
