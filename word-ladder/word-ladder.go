package leetcode

func ladderLength(beginWord, endWord string, words []string) int {
	dict := make(map[string]bool, len(words))
	for i := 0; i < len(words); i++ {
		dict[words[i]] = true
	}

	delete(dict, beginWord)
	queue := make([]string, 0, len(words))
	var trans func(string) bool

	trans = func(word string) bool {
		bytes := []byte(word)
		for i := 0; i < len(bytes); i++ {
			diffLetter := bytes[i]
			for j := 0; j < 26; j++ {
				b := 'a' + byte(j)
				if b == diffLetter {
					continue
				}
				bytes[i] = b
				if dict[string(bytes)] {
					// words 中存在 string(bytes)
					if string(bytes) == endWord {
						// trans 到了 endWord
						// 提前结束
						return true
					}

					// 把 string(bytes) 放入 queue 的尾部
					queue = append(queue, string(bytes))
					delete(dict, string(bytes))
				}
			}
			bytes[i] = diffLetter
		}
		return false
	}
	queue = append(queue, beginWord)
	dist := 1
	for len(queue) > 0 {
		qLen := len(queue)

		// 这个 for 循环，是按照每个 word 的 dist 值，来切分 queue 的
		for i := 0; i < qLen; i++ {
			// word 出列
			word := queue[0]
			queue = queue[1:]

			if trans(word) {
				// word 能够 trans 到 endWord
				// 提前结束
				return dist + 1
			}
		}

		dist++
	}

	return 0

}
