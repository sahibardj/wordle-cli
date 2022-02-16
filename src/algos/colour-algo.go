package algos

import "git.hanabi.in/dev/wordle-cli/src/colours"

type Lookup map[string]int

// Generate answer lookup table, and fill up with count of letters.
func GenAnsLookup(answer string) Lookup {
	ans_lookup := make(Lookup, 0)
	for _, ans_elem := range answer {
		ans_char := string(ans_elem)
		ans_lookup[ans_char]++
	}
	return ans_lookup
}

// Sahiba's algo -- commit your old code sk, I will patch the refactoring later -- ac.
// Get colours of the guess + the alphabet colouring too.
func GetColours(answer, guess string, ans_lookup, alphabet Lookup) string {
	// answer and guess required to check for green letters.
	var res string
	guess_lookup_table := make(Lookup, 0)
	for idx, guess_elem := range guess {
		guess_char := string(guess_elem)
		if ans_lookup[guess_char] == 0 { // Letter does not exist in answer, red.
			res += "R"
			alphabet[guess_char] = colours.Code_red
		} else { // If letter exists.
			guess_lookup_table[guess_char]++
			if guess_lookup_table[guess_char] <= ans_lookup[guess_char] {
				if answer[idx] == guess[idx] {
					res += "G"
					alphabet[guess_char] = colours.Code_green
				} else {
					res += "Y"
					if alphabet[guess_char] == colours.Code_grey { // Don't overwrite a Green alphabet with Yellow.
						alphabet[guess_char] = colours.Code_yellow
					}
				}
			} else { // Extra repeating characters.
				// alphabet colour would have been at least yellow, because it matched already, so don't edit.
				res += "R"
			}
		}
	}
	return res
}
