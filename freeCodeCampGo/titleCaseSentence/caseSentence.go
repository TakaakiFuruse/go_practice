package caseSentence

import (
	"strings"
	"unicode"
)

func titleCase(s string) string {
	stAr := strings.Split(s, " ")
	resAr := []string{}
	resStr := ""

	for _, value := range stAr {
		s := strings.ToLower(value)
		a := []rune(s)
		a[0] = unicode.ToUpper(a[0])
		resAr = append(resAr, string(a))
	}

	resStr = strings.Join(resAr, " ")
	return resStr
}
