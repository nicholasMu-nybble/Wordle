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

func getColorString(strout string, color string) string {
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
	return output
}

func printColorString(strout string, color string) {
	fmt.Print(getColorString(strout, color))
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
	// guess_words only contains words not in answer_words
	answerWordList := readWords("answer_words.txt")
	moreGuessWordList := readWords("guess_words.txt")
	input := ""
	wordHistory := ""
	guesses := 0
	allLetters := "qwertyuiopasdfghjklzxcvbnm"
	var accuracy [5]int
	var keyboardChecks [26]int
	for {
		// rand int for index and select it? be better in two lines maybe, but idc
		answerWord := answerWordList[rand.IntN(len(answerWordList))]
		guesses = 0
		wordHistory = ""
		clear(keyboardChecks[:])
		for ; guesses < 6; guesses++ {
			// get word
			for {
				// fmt.Println("Please enter a 5 letter word.")

				fmt.Print("\n\n\n\n\n\n\n\n")

				// print remaining letters
				for ind, letter := range allLetters {
					switch keyboardChecks[ind] {
					case 0:
						printColorString(string(letter), "white")
					case 1:
						printColorString(string(letter), "gray")
					case 2:
						printColorString(string(letter), "yellow")
					case 3:
						printColorString(string(letter), "green")
					}
					// fmt.Print(string(letter))
					if letter == 'p' || letter == 'l' || letter == 'm' {
						fmt.Print("\n")
					}
				}
				fmt.Print("\n")

				fmt.Print(wordHistory)

				fmt.Scanln(&input)
				input = strings.TrimSpace(input)
				// check length
				if len(input) != 5 {
					fmt.Println("Word is wrong length, try again")
					continue
				}
				// check validity
				if !slices.Contains(moreGuessWordList, input) && !slices.Contains(answerWordList, input) {
					fmt.Println(input, "is not a valid word. Try again.")
					continue
				}
				// if all's good, go next
				break
			}

			// check word correctness
			accuracy = getWordCorrectness(input, answerWord)
			for index, color := range accuracy {
				letterIndex := strings.Index(allLetters, string(input[index]))
				switch color {
				case 0:
					wordHistory += getColorString(string(input[index]), "gray")
					keyboardChecks[letterIndex] = max(keyboardChecks[letterIndex], 1)
				case 1:
					wordHistory += getColorString(string(input[index]), "yellow")
					keyboardChecks[letterIndex] = max(keyboardChecks[letterIndex], 2)
				case 2:
					wordHistory += getColorString(string(input[index]), "green")
					keyboardChecks[letterIndex] = max(keyboardChecks[letterIndex], 3)
				}
			}
			wordHistory += "\n"

			if input == answerWord {
				fmt.Println("Congratulations!")
				break
			}
		}

		fmt.Print("\n\n\n\n\n\n\n\n")

		fmt.Print(wordHistory)

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
