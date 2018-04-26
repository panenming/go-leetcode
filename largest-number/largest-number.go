package leetcode

import (
	"sort"
	"strconv"
)

type bytes [][]byte

func largestNumber(nums []int) string {
	size := len(nums)
	b := make(bytes, size)
	reSize := 0

	for i := range b {
		b[i] = []byte(strconv.Itoa(nums[i]))
		reSize += len(b[i])
	}

	sort.Sort(b)

	// 串联 b 成 res
	res := make([]byte, 0, reSize)
	for i := size - 1; i >= 0; i-- {
		res = append(res, b[i]...)
	}

	// 处理 res 以 0 开头的情况
	// 比如，[0, 0, 0] 的结果是 "000"
	i := 0
	for i < reSize-1 && res[i] == '0' {
		i++
	}

	return string(res[i:])
}

// ---- 实现sort需要的接口
func (b bytes) Len() int {
	return len(b)
}

func (b bytes) Less(i, j int) bool {
	size := len(b[i]) + len(b[j])

	bij := make([]byte, 0, size)
	bij = append(bij, b[i]...)
	bij = append(bij, b[j]...)

	bji := make([]byte, 0, size)
	bji = append(bji, b[j]...)
	bji = append(bji, b[i]...)

	for k := 0; k < size; k++ {
		if bij[k] < bji[k] {
			return true
		} else if bij[k] > bji[k] {
			return false
		}
	}

	return false
}

func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
