package main

import "fmt"

// Use the Parameter Testing feature in the box below to test your code with different arguments.
// Input = "coderbyte" Output = "etybredoc"
// Input = "I Love Code"Output = "edoC evoL I"

func Reverse(s string) string {
	words := []rune(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return string(words)
}

func main() {
	w := Reverse("eureka")
	fmt.Println(w)
}
