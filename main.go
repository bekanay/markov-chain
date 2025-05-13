package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalln("Incorrect usage")
	}

	cmd := flag.NewFlagSet("cmd", flag.ExitOnError)
	helpFlag := cmd.Bool("help", false, "Display usage instructions")
	wordNumber := cmd.Int("w", 100, "Number between 0 and 10000, default value equals to 100")
	prefix := cmd.String("p", "", "Starting prefix string (must exist in input)")
	prefixLength := cmd.Int("l", 2, "Prefix length between 1 and 5, default value equals to 2")
	cmd.Parse(args[1:])

	if *helpFlag {
		fmt.Println("Markov Chain text generator.\n")
		fmt.Println("Usage:")
		fmt.Println("  markovchain [-w <N>] [-p <S>] [-l <N>]")
		fmt.Println("  markovchain --help\n")
		fmt.Println("Options:")
		fmt.Println("  --help  Show this screen.")
		fmt.Println("  -w N    Number of maximum words")
		fmt.Println("  -p S    Starting prefix")
		fmt.Println("  -l N    Prefix length")
		os.Exit(0)
	}

	var words []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln("invalid text in stdin")
	}

	prefixFound := false

	wordsMap := make(map[string][]string)
	for i := 0; i+*prefixLength < len(words); i++ {
		prefixKey := strings.Join(words[i:i+*prefixLength], " ")
		suffix := words[i+*prefixLength]
		wordsMap[prefixKey] = append(wordsMap[prefixKey], suffix)
	}

	for i := 0; i < len(words); i++ {
		tempString := ""
		for j := 0; j < *prefixLength; j++ {
			if i+j >= len(words) {
				continue
			}
			tempString += words[i+j]
			if j != *prefixLength-1 {
				tempString += " "
			}
		}
		if *prefix == tempString {
			prefixFound = true
		}
	}

	tempPrefix := *prefix
	result := *prefix + " "
	for i := 0; i < *wordNumber; i++ {
		suffixes := wordsMap[tempPrefix]
		if len(suffixes) == 0 {
			break
		}
		suffix := suffixes[rand.Intn(len(suffixes))]
		result += suffix + " "
		tempPrefix = suffix
	}

	fmt.Println(result)

	if !prefixFound {
		log.Fatalln("prefix not found")
	}
}
