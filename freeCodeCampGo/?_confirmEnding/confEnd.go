package confEnd

import "strings"

func end(s ...string) bool {
	wordsToBeChecked := s[0]
	wordToCheck := s[1]
	wordsAr := strings.Split(wordsToBeChecked, " ")
	endLetter := wordsAr[len(wordsAr)-1][len(wordsAr[len(wordsAr)-1])-1]
	// endLetter := []byte{}

	// if len(wordsAr) == 1 && endLetter == wordToCheck {
	// 	return true
	// }

	if len(wordToCheck) > 1 && endLetter == wordToCheck[1] {
		return true
	}
	if wordsAr[len(wordsAr)-1] == wordToCheck {
		return true
	} else {
		return false
	}
}
