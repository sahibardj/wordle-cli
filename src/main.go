package main

import (
	"fmt"

	"git.hanabi.in/dev/wordle-cli/src/algos"
	"git.hanabi.in/dev/wordle-cli/src/utils"
)

const (
	chances   = 6
	word_size = 5
)

func main() {
	answer := utils.SelectAnswer()
	guesses := startGuessing(answer)
	fmt.Printf("Answer was: %s.\n", answer)
	utils.PrintShare(guesses)
}

func startGuessing(answer string) []string {

	var alphabet = utils.InitAlphabetTable()
	var anslookup = algos.GenAnsLookup(answer)
	var colouredChoices = []string{}
	fmt.Printf("Guess a %d-letter word.  You have %d tries.\n", word_size, chances)

	for cur_chance := 1; cur_chance <= chances; {
		utils.GuessPrompt(cur_chance)
		guess, err := utils.GetValidGuess(word_size)
		if err != nil {
			fmt.Printf("%v", err)
		} else {
			cur_chance++
			if guess != answer {
				colour_string := algos.GetColours(answer, guess, anslookup, alphabet)
				colouredChoices = append(colouredChoices, colour_string)
				utils.PrintColouredGuess(colour_string, guess)
			} else if guess == answer {
				colour_string := "GGGGG" // If the answer was correct, GetColours is not called, hence hard-coding.
				colouredChoices = append(colouredChoices, colour_string)
				fmt.Println("Correct guess!")
				break
			}
		}
		utils.PrintColouredAlpha(alphabet)
	}
	return colouredChoices
}
