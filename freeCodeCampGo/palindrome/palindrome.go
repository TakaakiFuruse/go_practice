package palindrome

import (
	"regexp"
	"strings"
)

func palindrome(s string) bool {
	rep := regexp.MustCompile(`[\W_]`)
	repStr := rep.ReplaceAllString(s, "")
	strAr := strings.Split(repStr, "")
	revAr := strAr

	for i, j := 0, len(revAr)-1; i < j; i, j = i+1, j-1 {
		revAr[i], revAr[j] = revAr[j], revAr[i]
	}

	jStr := strings.Join(revAr, "")
	return strings.ToLower(repStr) == strings.ToLower(jStr)
}
