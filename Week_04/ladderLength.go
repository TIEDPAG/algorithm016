package t127

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}

	g := makeGraph(wordList)

	return bfs(beginWord, endWord, g)
}

func bfs(beginWord, endWord string, graph map[string][]string) int {
	flag := make(map[string]bool)
	flag[beginWord] = true

	l := len(beginWord)

	queue := []struct {
		word  string
		level int
	}{{
		word:  beginWord,
		level: 1,
	}}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		for i := 0; i < l; i++ {
			key := fmt.Sprintf("%s-%s", node.word[:i], node.word[i+1:])
			for _, word := range graph[key] {
				if word == endWord { // 广度优先搜索  先找到则是最浅的路径
					return node.level + 1
				}

				if !flag[word] {
					flag[word] = true
					queue = append(queue, struct {
						word  string
						level int
					}{word: word, level: node.level + 1})
				}
			}
		}
	}
	return 0
}

func makeGraph(wordList []string) map[string][]string {
	l := len(wordList[0])
	rs := make(map[string][]string)
	for _, word := range wordList {
		for i := 0; i < l; i++ {
			key := fmt.Sprintf("%s-%s", word[:i], word[i+1:])
			arr, has := rs[key]
			if has {
				rs[key] = append(arr, word)
				continue
			}
			rs[key] = []string{word}
		}
	}
	return rs
}
