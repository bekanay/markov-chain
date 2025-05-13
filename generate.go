package main

import (
	"math/rand"
	"strings"
)

func Generate(wordsMap Model, start string, totalWords int, prefixLength int) string {
	tempPrefix := start
	result := start + " "
	for i := 0; i < totalWords-prefixLength; i++ {
		suffixes := wordsMap[tempPrefix]
		if len(suffixes) == 0 {
			break
		}
		suffix := suffixes[rand.Intn(len(suffixes))]
		result += suffix + " "
		parts := strings.Split(tempPrefix, " ")
		parts = append(parts[1:], suffix)
		tempPrefix = strings.Join(parts, " ")
	}
	return result
}
