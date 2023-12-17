package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(input string) []string {
	topWordsNumber := 10
	frequencyMap := make(map[string]int)
	words := strings.Fields(input)

	for _, word := range words {
		frequencyMap[word]++
	}

	keys := make([]string, 0, len(frequencyMap))
	for key := range frequencyMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if frequencyMap[keys[i]] == frequencyMap[keys[j]] {
			switch strings.Compare(keys[i], keys[j]) {
			case -1:
				return true
			case 0, 1:
				return false
			}
		}

		return frequencyMap[keys[i]] > frequencyMap[keys[j]]
	})

	top := make([]string, 0, topWordsNumber)
	keysLen := len(keys)
	for i := 0; i < topWordsNumber && i < keysLen; i++ {
		top = append(top, keys[i])
	}

	return top
}
