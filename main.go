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
	cmd := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	helpFlag := cmd.Bool("help", false, "Display usage instructions")
	wordNumber := cmd.Int("w", 100, "Number between 0 and 10000, default value equals to 100")
	prefix := cmd.String("p", "", "Starting prefix string (must exist in input)")
	prefixLength := cmd.Int("l", 2, "Prefix length between 1 and 5, default value equals to 2")
	cmd.Parse(os.Args[1:])

	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	if *wordNumber == 0 {
		os.Exit(0)
	}

	if *wordNumber < 0 || *wordNumber > 10000 {
		log.Fatalln("invalid number of words")
	}

	if *prefixLength < 1 || *prefixLength > 5 {
		log.Fatalln("invalid prefix length")
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

	if len(words) < *prefixLength {
		log.Fatalln("input length is less than prefix length")
	}

	if *prefix == "" {
		*prefix = strings.Join(words[0:*prefixLength], " ")
	}

	prefixSlice := strings.Split(*prefix, " ")
	if len(prefixSlice) != *prefixLength {
		log.Fatalln("prefix length shoud be equal to set length")
	}

	if *wordNumber <= len(prefixSlice) {
		fmt.Println(strings.Join(prefixSlice[:*wordNumber], " "))
		os.Exit(0)
	}

	wordsMap := make(map[string][]string)
	for i := 0; i+*prefixLength <= len(words); i++ {
		prefixKey := strings.Join(words[i:i+*prefixLength], " ")
		if i+*prefixLength < len(words) {
			wordsMap[prefixKey] = append(wordsMap[prefixKey], words[i+*prefixLength])
		} else {
			if _, ok := wordsMap[prefixKey]; !ok {
				wordsMap[prefixKey] = []string{}
			}
		}
	}

	suffixes := wordsMap[*prefix]
	if suffixes == nil {
		log.Fatalln("prefix not found")
	}

	if len(suffixes) == 0 {
		fmt.Println(*prefix)
		os.Exit(0)
	}

	// rand.Seed(time.Now().UnixNano())

	tempPrefix := *prefix
	result := *prefix + " "
	for i := 0; i < *wordNumber-len(prefixSlice); i++ {
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

	fmt.Println(result)
	os.Exit(0)
}

func printHelp() {
	fmt.Fprintf(os.Stdout, `Markov Chain text generator.

Usage:
markovchain [-w <N>] [-p <S>] [-l <N>]
markovchain --help

Options:
--help  Show this screen.
-w N    Number of maximum words
-p S    Starting prefix
-l N    Prefix length
`)
}
