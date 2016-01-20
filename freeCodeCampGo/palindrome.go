package main

import (
	"fmt"
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

func main() {
	fmt.Println(palindrome("eye") == true)
	fmt.Println(palindrome("race car") == true)
	fmt.Println(palindrome("not a palindrome") == false)
	fmt.Println(palindrome("A man, a plan, a canal. Panama") == true)
	fmt.Println(palindrome("never odd or even") == true)
	fmt.Println(palindrome("nope") == false)
	fmt.Println(palindrome("almostomla") == false)
	fmt.Println(palindrome("My age is 0, 0 si ega ym.") == true)
	fmt.Println(palindrome("1 eye for of 1 eye.") == false)
}
