package confEnd

import "strings"

func end(s ...string) bool {
	wordToCheck := s[1]
	wordToCheckAr := strings.Split(wordToCheck, "")
	wordsToBeChecked := s[0]
	wordsToBeCheckedAr := strings.Split(wordsToBeChecked, " ")
	lasLetterOfWord := strings.Split(wordsToBeCheckedAr[len(wordsToBeCheckedAr)-1], "")
	endLetter := lasLetterOfWord[len(lasLetterOfWord)-1]

	if len(wordsToBeCheckedAr) == 1 && endLetter == wordToCheck {
		return true
	}

	if wordsToBeCheckedAr[len(wordsToBeCheckedAr)-1] == wordToCheck {
		return true
	} else {
		lw := strings.Split(wordsToBeCheckedAr[len(wordsToBeCheckedAr)-1], "")
		for i, j := 0, len(lw)-1; i < j; i, j = i+1, j-1 {
			lw[i], lw[j] = lw[j], lw[i]
		}
		for i, j := 0, len(wordToCheckAr)-1; i < j; i, j = i+1, j-1 {
			wordToCheckAr[i], wordToCheckAr[j] = wordToCheckAr[j], wordToCheckAr[i]
		}

		for i := 0; i < len(wordToCheckAr); i++ {
			if wordToCheckAr[i] != lw[i] {
				return false
			}
		}
		return true
	}
}
