package main

import (
	"fmt"
	"os"
)

func readWords(path string) []string {
	file, err := os.Open(path) // open file
	if err != nil { // if error, return nil, print error
		fmt.Println("Error opening file:", err)
		return nil
	}

	fmt.Println("File opened:", file.Name()) // file objects have names
	return nil // temp
}

func main() {
	readWords("words.txt")
	readWords("wordsa.txt") // file fails to open
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