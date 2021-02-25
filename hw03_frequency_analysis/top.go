package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

type TopWord struct {
	count int
	word  string
}

// Top10 is getting top 10 words from text.
func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}
	textSlice := strings.FieldsFunc(text, unicode.IsSpace)

	topWords := map[string]int{}
	for _, s := range textSlice {
		s = strings.TrimSpace(s)
		_, ok := topWords[s]
		if ok {
			topWords[s]++
		} else {
			topWords[s] = 1
		}
	}

	topWordSlice := make([]TopWord, 0, len(topWords))
	for k, v := range topWords {
		topWordSlice = append(topWordSlice, TopWord{count: v, word: k})
	}

	sort.SliceStable(topWordSlice, func(i, j int) bool {
		if topWordSlice[i].count == topWordSlice[j].count {
			return topWordSlice[i].word < topWordSlice[j].word
		}
		return topWordSlice[i].count > topWordSlice[j].count
	})
	res := make([]string, 0, 10)
	for i := 0; i <= 9 && i <= len(topWordSlice)-1; i++ {
		res = append(res, topWordSlice[i].word)
	}
	return res
}
