package main

import (
	"fmt"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordLen := len(beginWord)
	hash := make(map[string]map[string]bool)
	for _, word := range wordList {
		for i := 0; i < wordLen; i++ {
			temp := word[:i] + "*" + word[i+1:]
			if _, ok := hash[temp]; !ok {
				hash[temp] = make(map[string]bool)
			}
			hash[temp][word] = true
		}
	}

	res := make([][]string, 0)
	visited := make(map[string]bool)
	queue := [][]string{{beginWord}}
	for len(queue) > 0 {
		if len(res) > 0 {
			break
		}

		queLen := len(queue)
		subVisited := make(map[string]bool)
		for i := 0; i < queLen; i++ {
			path := queue[0]
			queue = queue[1:]
			word := path[len(path)-1]

			for j := 0; j < wordLen; j++ {
				temp := word[:j] + "*" + word[j+1:]
				targets, ok := hash[temp]
				if !ok {
					continue
				}
				for target := range targets {
					path = append(path, target)
					if target == endWord {
						res = append(res, append([]string{}, path...))
						continue
					}

					if !visited[target] {
						queue = append(queue, append([]string{}, path...))
					}

					subVisited[target] = true
					path = path[:len(path)-1]
				}
			}

			path = path[:len(path)-1]
		}

		for word := range subVisited {
			visited[word] = true
		}
	}
	return res
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}
