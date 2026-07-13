package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func readWords(path string) []string {
	file, err := os.Open(path) // open file
	if err != nil { // if error, return nil, print error
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close() // close file AFTER function returns

	data, err := io.ReadAll(file) // read the file into data
	if err != nil { // error handling again
		fmt.Println("Error opening file:", err)
		return nil
	}

	// parse data
	output := strings.Split(string(data), "\n")
	return output
}

func main() {
	wordList := readWords("words.txt")
	for _, value := range wordList {
		fmt.Println(value)
	}
	fmt.Println(wordList[0])
}

/*
randomly select a 5 letter word
	from words.txt
	gotta find out file parsing and rng in go
	save list of words in array, use for word checking too

input 5 letter word
	how take input? idk figure it out
	early test: text boxes, check length
	by end: input letters into boxes, preventing 6+ letters entirely
		will need ui somehow idk how on go, but i'll figure it out
	check that it's actually a word

display word

show status of each letter
	check for greens, then yellows, anything else is gray
	temporary array, remove letters when found to avoid duplicates when no duplicates

track status of each letter
	array of ints? 0 unchecked, 1 gray, 2 yellow, 3 green?

limit to 6 words, then show word
	track turn number
*/