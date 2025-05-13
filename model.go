package main

import "strings"

type Model map[string][]string

func BuildModel(words []string, prefixLength int) Model {
	wordsMap := make(map[string][]string)
	for i := 0; i+prefixLength <= len(words); i++ {
		prefixKey := strings.Join(words[i:i+prefixLength], " ")
		if i+prefixLength < len(words) {
			wordsMap[prefixKey] = append(wordsMap[prefixKey], words[i+prefixLength])
		} else {
			if _, ok := wordsMap[prefixKey]; !ok {
				wordsMap[prefixKey] = []string{}
			}
		}
	}
	return wordsMap
}
