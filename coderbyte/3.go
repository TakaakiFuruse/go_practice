package main

// Using the JavaScript language, have the function LongestWord(sen) take the sen parameter being passed and return the largest word in the string. If there are two or more words that are the same length, return the first word from the string with that length. Ignore punctuation and assume sen will not be empty.
// Input = "fun&!! time"Output = "time"
// Input = "I love dogs"Output = "love"

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	word := strings.Fields("take the sen parameter being passed and return the largest")
	sort.Strings(word)
	fmt.Println(word)

}
