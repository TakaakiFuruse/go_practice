package main

import (
	"fmt"
	"strings"
)

func main() {
	for i, j := 1, 3; i < 9 && j > -1; i, j = i+2, j-1 {
		fmt.Println(strings.Repeat(" ", j) + strings.Repeat("*", i))

	}
	for j, h := 5, 1; j > 0 && h < 4; j, h = j-2, h+1 {
		fmt.Println(strings.Repeat(" ", h) + strings.Repeat("*", j))
	}
}
