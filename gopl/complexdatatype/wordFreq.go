package main

import (
	"unicode"
)

func wordFreq(s string) (m map[string]int) {
	m = make(map[string]int)
	for _, value := range s {
		if unicode.IsNumber(value) {
			m["number"] += 1
		}
		if unicode.IsLower(value) {
			m["lower"] += 1
		}
		if unicode.IsUpper(value) {
			m["upper"] += 1
		}
	}
	return m
}

/*func main() {
	sentences := "dsjgfdhdb4hguebdrg7f%$Edcydvfe7cdfvw76fvfdhvbbu&T&Fvydgr6hefxfbvud*IUJIUHI&U"
	fmt.Println(wordFreq(sentences))
}
*/