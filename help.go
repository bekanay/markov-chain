package main

import (
	"fmt"
	"os"
)

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
