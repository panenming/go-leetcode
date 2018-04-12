package leetcode

func sortColors(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	temp := nums[0]
	nums[0] = 1
	// 三路快排
	// 排序完成后，
	//    i 指向 0 后的 1，
	//    j 指向 最后一个 1 后面的位置，
	//    k 指向 2 前面的 1，
	// 在整个排序过程中，nums[i:j]中始终都是1
	i, j, k := 0, 1, length-1
	for j <= k {
		switch {
		case nums[j] < 1:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1 < nums[j]:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		default:
			j++
		}
	}
	switch temp {
	case 0:
		nums[i] = temp
	case 2:
		nums[k] = temp
	}
	return
}
