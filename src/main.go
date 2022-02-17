package main

import (
	gameplay "git.hanabi.in/dev/wordle-cli/src/utils"
)

func main() {
	answer := gameplay.SelectAnswer()
	guesses := gameplay.StartGuessing(answer)
	gameplay.GracefullyFinishGame(answer, guesses)
}
