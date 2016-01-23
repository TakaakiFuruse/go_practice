package confEnd

import "strings"

func end(s ...string) bool {
	wordsToBeChecked := s[0]
	wordToCheck := s[1]
	wordsAr := strings.Split(wordsToBeChecked, " ")
	lasLetterOfWord := strings.Split(wordsAr[len(wordsAr)-1], "")
	endLetter := lasLetterOfWord[len(lasLetterOfWord)-1]

	if len(wordsAr) == 1 && endLetter == wordToCheck {
		return true
	}

	if len(wordToCheck) > 1 && endLetter == wordToCheck {
		return true
	}
	if wordsAr[len(wordsAr)-1] == wordToCheck {
		return true
	} else {
		return false
	}
}
