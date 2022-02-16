package algos

type Lookup map[string]int

func GetColours(ans, guess string, alphabet map[string]int) string {
	alookup := getALookup(ans)
	res := glookup(ans, guess, alookup, alphabet)
	return res
}

func getALookup(ans string) Lookup {
	alookup := make(Lookup, 0)
	alen := len(ans)
	for i := 0; i < alen; i++ {
		achar := string(ans[i])
		alookup[achar]++
	}
	return alookup
}

// Sahiba's algo
func glookup(ans, guess string, alookup Lookup, alphabet map[string]int) string {
	var res string
	glookup := make(Lookup, 0)
	glen := len(guess)
	for i := 0; i < glen; i++ {
		gchar := string(guess[i])
		if alookup[gchar] == 0 {
			res += "R"
			alphabet[gchar] = -1
		} else {
			glookup[gchar]++
			if glookup[gchar] <= alookup[gchar] {
				if ans[i] == guess[i] {
					res += "G"
					alphabet[gchar] = 1
				} else {
					res += "Y"
					alphabet[gchar] = 0
				}
			} else {
				res += "R"
			}
		}
	}
	return res
}
