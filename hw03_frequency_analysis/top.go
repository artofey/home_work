package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

type TopWord struct {
	count int
	word  string
}

// Top10 is getting top 10 words from text.
func Top10(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	textSlice := strings.Split(text, "\t")
	var b strings.Builder
	for _, item := range textSlice {
		b.WriteString(item)
		b.WriteString(" ")
	}
	text = b.String()
	textSlice = strings.Split(text, " ")
	sort.SliceStable(textSlice, func(i, j int) bool { return textSlice[i] < textSlice[j] })
	topWordSlice := []TopWord{}
	var sLeft string
	for _, s := range textSlice {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		if sLeft == s {
			topWordSlice[len(topWordSlice)-1].count++
		} else {
			topWordSlice = append(topWordSlice, TopWord{1, s})
		}
		sLeft = s
	}
	sort.SliceStable(topWordSlice, func(i, j int) bool { return topWordSlice[i].count > topWordSlice[j].count })
	res := []string{}
	for i := 0; i <= 9; i++ {
		if i > len(topWordSlice)-1 {
			break
		}
		res = append(res, topWordSlice[i].word)
	}
	return res
}
