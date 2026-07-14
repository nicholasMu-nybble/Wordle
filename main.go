package main

import (
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

func readWords(path string) []string {
	file, err := os.Open(path) // open file
	if err != nil {            // if error, return nil, print error
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close() // close file AFTER function returns

	data, err := io.ReadAll(file) // read the file into data
	if err != nil {               // error handling again
		fmt.Println("Error opening file:", err)
		return nil
	}

	// parse data
	output := strings.Split(string(data), "\n")
	for index := range output {
		// removes the \r from everything
		output[index] = strings.TrimSpace(output[index])
	}
	return output
}

func printColorString(strout string, color string) {
	output := ""
	color = strings.ToLower(color)
	switch color {
		case "red":
			output += "\033[31m"
		case "yellow":
			output += "\033[33m"
		case "green":
			output += "\033[32m"
		case "gray", "grey":
			output += "\033[90m"
		case "white":
			output += "\033[97m"
	}
	output += strout
	output += "\033[0m"
	fmt.Print(output)
}

func getWordCorrectness(guess string, correct string) [5]int {
	// 2: green, 1: yellow, 0: gray
	guess_letters := strings.Split(guess, "")
	correct_letters := strings.Split(correct, "")
	output := [5]int{0, 0, 0, 0, 0}
	for index := range guess_letters {
		// green O(n) (5 checks)
		if guess_letters[index] == correct_letters[index] {
			output[index] = 2
			guess_letters[index] = ""
			correct_letters[index] = " " // different as to not confuse yellow part
		}
	}

	// yellow O(n^2) (25 checks)
	for g := range guess_letters {
		for c := range correct_letters {
			if guess_letters[g] == correct_letters[c] {
				output[g] = 1
				guess_letters[g] = ""
				correct_letters[c] = " "
			}
		}
	}

	// anything left will be gray

	return output
}

func main() {
	wordList := readWords("words.txt")
	input := ""
	guesses := 0
	var accuracy [5]int
	for {
		// rand int for index and select it? be better in two lines maybe, but idc
		answerWord := wordList[rand.IntN(len(wordList))]
		guesses = 0
		for ; guesses < 6; guesses++ {
			// get word
			for {
				fmt.Println("Please enter a 5 letter word.")
				fmt.Scanln(&input)
				input = strings.TrimSpace(input)
				// check length
				if len(input) != 5 {
					fmt.Println("Word is wrong length, try again")
					continue
				}
				// check validity
				if !slices.Contains(wordList, input) {
					fmt.Println(input, "is not a valid word. Try again.")
					continue
				}
				// if all's good, go next
				break
			}

			// check word correctness
			accuracy = getWordCorrectness(input, answerWord)
			for index, color := range accuracy {
				switch color {
					case 0:
						printColorString(string(input[index]), "gray")
					case 1:
						printColorString(string(input[index]), "yellow")
					case 2:
						printColorString(string(input[index]), "green")
				}
			}
			fmt.Println("")


			if input == answerWord {
				fmt.Println("Congratulations!")
				break
			}
		}

		// game over
		fmt.Println("The answer was", answerWord)
		for {
			fmt.Println("Would you like to play again? (y/n)")
			fmt.Scanln(&input) // interesting, pass a pointer to the output string
			if input == "y" || input == "n" {
				break
			}
			fmt.Println("Invalid input, try again.")
		}
		if input == "n" {
			break
		}
	}
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
