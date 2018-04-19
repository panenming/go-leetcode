package leetcode

func findLadders(beginWord, endWord string, words []string) [][]string {
	words = deleteBeginWord(words, beginWord)
	trans := map[string][]string{}
	isTransedEndWord := false
	cnt := 1
	var bfs func([]string, []string)
	bfs = func(words, nodes []string) {
		cnt++
		newWords := make([]string, 0, len(words))
		newNodes := make([]string, 0, len(words))
		for _, w := range words {
			isTransed := false
			for _, n := range nodes {
				if isTransable(n, w) {
					trans[n] = append(trans[n], w)
					isTransed = true
				}
			}

			if isTransed {
				newNodes = append(newNodes, w)
				if w == endWord {
					isTransedEndWord = true
				}
			} else {
				newWords = append(newWords, w)
			}
		}
		if isTransedEndWord || // 转换到了 endWord 说明已经找到了所有的最短路径
			len(newWords) == 0 || // words 的所有单词都已经可以从 beginWord trans 到
			len(newNodes) == 0 { // newWords 中单词，是 beginWord 无法 trans 到的
			return
		}
		bfs(newWords, newNodes)
	}

	nodes := []string{beginWord}
	bfs(words, nodes)
	res := [][]string{}
	if !isTransedEndWord {
		// beginWord 无法 trans 到 endWord
		return res
	}

	path := make([]string, cnt)
	path[0] = beginWord

	var dfs func(int)
	// 使用 dfs 方法，生成最短路径
	dfs = func(idx int) {
		if idx == cnt {
			// path 已经填充完毕
			if path[idx-1] == endWord {
				// 最后一个单词是 endWord，说明这是一条最短路径
				res = append(res, deepCopy(path))
			}
			return
		}

		prev := path[idx-1]
		for _, w := range trans[prev] {
			// 利用 prev->w 填充 path[idx]
			path[idx] = w
			dfs(idx + 1)
		}
	}

	dfs(1)

	return res
}

func deepCopy(src []string) []string {
	temp := make([]string, len(src))
	copy(temp, src)
	return temp
}

func deleteBeginWord(words []string, beginWord string) []string {
	i, size := 0, len(words)
	for ; i < size; i++ {
		if words[i] == beginWord {
			break
		}
	}

	if i == size {
		return words
	}

	words[i] = words[size-1]
	return words[:size-1]
}

func isTransable(a, b string) bool {
	// onceAgain == true 说明已经出现过不同的字符了
	onceAgain := false
	for i := range a {
		if a[i] != b[i] {
			if onceAgain {
				return false
			}
			onceAgain = true
		}
	}
	return true
}
