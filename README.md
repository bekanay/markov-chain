# markov-chain

A simple text generator based on the **Markov Chain algorithm**, written in Go.  
It reads an input text and generates statistically coherent output by learning prefix-suffix word patterns.

## 📚 Features

- Reads input from standard input (`stdin`)
- Supports:
  - custom **prefix length** (`-l`)
  - custom **starting prefix** (`-p`)
  - custom **word limit** (`-w`)
- Prints natural-sounding random text
- Handles errors and edge cases cleanly

## 📌 Usage

```sh
# Build
go build -o markovchain .

# Basic usage (generates 100 words by default)
cat input.txt | ./markovchain

# Custom max words
cat input.txt | ./markovchain -w 50

# Custom starting prefix
cat input.txt | ./markovchain -p "the world"

# Custom prefix length
cat input.txt | ./markovchain -l 3 -p "in the beginning"

# Show help
./markovchain --help
