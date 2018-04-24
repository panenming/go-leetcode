package leetcode
func findMin(a []int) int{
	len := len(a)
	i := 1
	for i < len && a[i-1] < a[i]{
		i++
	}
	return a[i%len]
}