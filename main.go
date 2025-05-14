package main

import (
	"flag"
	"fmt"
	"log"
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

	words, err := ReadWords(os.Stdin)
	if err != nil {
		log.Fatalln("error reading stdin: ", err)
	}

	if len(words) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no input text")
		os.Exit(1)
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

	wordsMap := BuildModel(words, *prefixLength)

	suffixes := wordsMap[*prefix]
	if suffixes == nil {
		log.Fatalln("prefix not found")
	}

	if len(suffixes) == 0 {
		fmt.Println(*prefix)
		os.Exit(0)
	}

	result := Generate(wordsMap, *prefix, *wordNumber, *prefixLength)

	fmt.Println(result)
}
