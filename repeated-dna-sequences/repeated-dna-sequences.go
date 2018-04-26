package leetcode

import (
	"sort"
)

func findRepeatedDnaSequences(s string) []string {
	var res []string

	if len(s) <= 10 {
		return nil
	}

	in := make(map[string]bool, len(s)-9)
	for i := 0; i+10 <= len(s); i++ {
		str := s[i : i+10]
		if _, ok := in[str]; ok {
			res = append(res, str)
		}

		in[str] = true
	}

	sort.Strings(res)
	return res
}
