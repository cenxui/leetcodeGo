package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "cars"
	wordDict := []string{"car","ca","rs"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	return wordBreakSet(s, 0, wordDict, make(map[int]interface{}))
}

func wordBreakSet(s string, p int, wordDict []string, cache map[int]interface{}) bool {
	if len(s) == p {
		return true
	}

	if _, ok := cache[p]; ok {
		return false
	}

	for _, w :=range wordDict {
		if strings.HasPrefix(s[p:], w) {
			if wordBreakSet(s, p + len(w) , wordDict, cache) {
				return true
			}else {
				cache[p + len(w)] = nil
			}
		}
	}
	return false
}