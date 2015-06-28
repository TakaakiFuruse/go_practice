package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

var (
	i, j      int = 1, 2
	arr_key       = []string{"i", "j", "one", "two", "three"}
	arr_value     = []string{"a", "a", "a", "b", "b"}
)

func main() {
	for i := 0; i < 6; i++ {
		if arr_value[i] == "a" {
			fmt.Println(arr_key[i], "is", arr_value[i])
		} else {
			fmt.Println(arr_key[i], "??? that's not a!!")
		}
	}
}
