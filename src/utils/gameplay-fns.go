package gameplay

import (
	"fmt"
	"math/rand"
	"strings"

	"git.hanabi.in/dev/wordle-cli/src/algos"
	"git.hanabi.in/dev/wordle-cli/src/colours"
	"git.hanabi.in/dev/wordle-cli/src/data"
)

const (
	chances   = 6
	word_size = 5
)

// From a pre-defined sorted list of words, pick a random word which is the answer.
func SelectAnswer() string {
	words := []string{}
	for _, elem := range data.Answers {
		words = append(words, elem)
	}
	algos.Shuffle(words)
	index := rand.Intn(len(words))
	return words[index]
}

// Create a lookup table for alphabet, initialise them to grey colour.
func initAlphabetTable() algos.Lookup {
	var alphabet_table = make(algos.Lookup, 0)
	letters := "abcdefghijklmnopqrstuvwxyz"
	for _, elem := range letters {
		letter_char := string(elem)
		alphabet_table[letter_char] = colours.Code_grey
	}
	return alphabet_table
}

// Check if the word is of the size of answer.  Returns true if size mismatch.
func isWrongGuessSize(guess string) bool {
	return len(guess) != word_size
}

// Get user input for the guess, and process it (lower case)
func fetchGuess() string {
	var guess string
	fmt.Scanf("%s", &guess)
	guess = strings.ToLower(guess)
	return guess
}

// Check if guessed word is not in the list of possible words.  Returns true if not found.
func isNotValidWord(guess string) bool {
	idx1 := algos.BinarySearch(data.Answers, guess)
	idx2 := algos.BinarySearch(data.Guesses, guess)
	return idx1 == -1 && idx2 == -1
}

// returns either a valid guess, XOR an error.
func getValidGuess(prev_guesses []string) (string, error) {
	guess := fetchGuess()
	var error_msg error = nil
	if isWrongGuessSize(guess) {
		error_msg = fmt.Errorf("Word should be of length %d.\n", word_size)
	} else if isNotValidWord(guess) {
		error_msg = fmt.Errorf("Not a valid word.\n")
	} else if isOldGuess(guess, prev_guesses) {
		error_msg = fmt.Errorf("You already guessed this word.\n")
	}
	return guess, error_msg
}

// Check if the current guess was already guessed or not.
func isOldGuess(guess string, prev_guesses []string) bool {
	for _, elem := range prev_guesses {
		if elem == guess {
			return true
		}
	}
	return false
}

// Print the guess prompt.
func guessPrompt(chance int) {
	msg := colours.Bold(fmt.Sprintf("Guess #%d?: ", chance))
	fmt.Print(msg)
}

// Look at colour_string (Y,G,R) and print the characters of word in colour.
func printColouredGuess(colour_string, word string) {
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
func printColouredAlpha(alphabet algos.Lookup) {
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
func printShare(guesses []string) {
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

// Prompt if the answer should be printed.
func shouldPrintAnswer() bool {
	var ans string
	fmt.Print(colours.Bold("Show answer?[yN]: "))
	fmt.Scanf("%s", &ans)
	if ans == "Y" || ans == "y" {
		return true
	}
	return false
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

// The guessing function, returns guess history and if the user won the game.
func StartGuessing(answer string) ([]string, bool) {

	var alphabet = initAlphabetTable()
	var anslookup = algos.GenAnsLookup(answer)
	var prev_guesses = []string{}
	var colouredChoices = []string{}
	didWin := false
	fmt.Printf("Guess a %d-letter word.  You have %d tries.\n", word_size, chances)

	for cur_chance := 1; cur_chance <= chances; {
		guessPrompt(cur_chance)
		guess, err := getValidGuess(prev_guesses)
		if err != nil {
			fmt.Printf("%v", err)
		} else {
			cur_chance++
			prev_guesses = append(prev_guesses, guess)
			colour_string := algos.GetColours(answer, guess, anslookup, alphabet)
			colouredChoices = append(colouredChoices, colour_string)
			printColouredGuess(colour_string, guess)
			if guess == answer {
				didWin = true
				break
			}
		}
		printColouredAlpha(alphabet)
	}
	return colouredChoices, didWin
}

// Handle end of the game once correct answer is reached, or when all chances are over.
func GracefullyFinishGame(answer string, guesses []string, didWin bool) {
	if !didWin && shouldPrintAnswer() {
		fmt.Printf("Answer was: %s.\n", answer)
	}
	printShare(guesses)
}
