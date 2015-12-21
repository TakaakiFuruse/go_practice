package main

import "fmt"

// calculate factorial

func factorial(n int) int {

	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
	return n
}

func main() {
	fmt.Println(factorial(20))
}
