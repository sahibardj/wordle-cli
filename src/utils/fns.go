package utils

import (
	"fmt"
	"math/rand"
	"strings"

	"git.hanabi.in/dev/wordle-cli/src/algos"
	"git.hanabi.in/dev/wordle-cli/src/colours"
	"git.hanabi.in/dev/wordle-cli/src/data"
)

// From a pre-defined sorted list of words, pick a random word which is the answer.
func SelectAnswer() string {
	words := []string{}
	for _, elem := range data.Words {
		words = append(words, elem)
	}
	algos.Shuffle(words)
	index := rand.Intn(len(words))
	return words[index]
}

// Create a lookup table for alphabet, initialise them to grey colour.
func InitAlphabetTable() algos.Lookup {
	var alphabet_table = make(algos.Lookup, 0)
	letters := "abcdefghijklmnopqrstuvwxyz"
	for _, elem := range letters {
		letter_char := string(elem)
		alphabet_table[letter_char] = colours.Code_grey
	}
	return alphabet_table
}

// Check if the word is of the size of answer.  Returns true if size mismatch.
func IsWrongGuessSize(guess string, word_size int) bool {
	return len(guess) != word_size
}

// Get user input for the guess, and process it (lower case)
func FetchGuess() string {
	var guess string
	fmt.Scanf("%s", &guess)
	guess = strings.ToLower(guess)
	return guess
}

// Check if guessed word is not in the list of possible words.  Returns true if not found.
func IsNotValidWord(guess string) bool {
	idx := algos.BinarySearch(data.Words, guess)
	return idx == -1
}

// returns either a valid guess, XOR an error.
func GetValidGuess(word_size int) (string, error) {
	guess := FetchGuess()
	if IsWrongGuessSize(guess, word_size) {
		error_msg := fmt.Errorf("Word should be of length %d.\n", word_size)
		return guess, error_msg
	}
	if IsNotValidWord(guess) {
		error_msg := fmt.Errorf("Not a valid word.\n")
		return guess, error_msg
	}
	return guess, nil
}

// Print the guess prompt.
func GuessPrompt(chance int) {
	msg := colours.Bold(fmt.Sprintf("Guess #%d?: ", chance))
	fmt.Print(msg)
}

// Look at colour_string (Y,G,R) and print the characters of word in colour.
func PrintColouredGuess(colour_string, word string) {
	for idx, word_elem := range word {
		col_str_char := string(colour_string[idx])
		word_char := string(word_elem)
		if col_str_char == "R" {
			fmt.Print(colours.Red(word_char))
		} else if col_str_char == "Y" {
			fmt.Print(colours.Yellow(word_char))
		} else if col_str_char == "G" {
			fmt.Print(colours.Green(word_char))
		}
	}
	fmt.Println()
}

// Print coloured alphabet for aiding which words to guess.
func PrintColouredAlpha(alphabet algos.Lookup) {
	kbd_rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"} // keyboard layout printing.
	for _, kbd_row := range kbd_rows {
		for _, kbd_key := range kbd_row {
			kbd_key_char := string(kbd_key)
			char_col_code := alphabet[kbd_key_char]
			if char_col_code == colours.Code_grey {
				fmt.Print(colours.Grey(kbd_key_char))
			} else if char_col_code == colours.Code_red {
				fmt.Print(colours.Red(kbd_key_char))
			} else if char_col_code == colours.Code_yellow {
				fmt.Print(colours.Yellow(kbd_key_char))
			} else if char_col_code == colours.Code_green {
				fmt.Print(colours.Green(kbd_key_char))
			}
			fmt.Print("  ")
		}
		fmt.Println()
	}
	fmt.Println()
}

// Print share emoji.
func PrintShare(guesses []string) {
	if shouldPrintShareEmojis() {
		fmt.Println()
		for _, row := range guesses {
			for _, elem_byte := range row {
				elem := string(elem_byte)
				if elem == "R" {
					fmt.Print("🌑")
				} else if elem == "Y" {
					fmt.Print("🌕")
				} else if elem == "G" {
					fmt.Print("✅")
				}
			}
			fmt.Println()
		}
	}
}

// Prompt if the share emojies be printed.
func shouldPrintShareEmojis() bool {
	var ans string
	fmt.Print(colours.Bold("Share your results?[yN]: "))
	fmt.Scanf("%s", &ans)
	if ans == "Y" || ans == "y" {
		return true
	}
	return false
}
