package main

import (
	"fmt"
)

var (
	arr_key   = []string{"i", "j", "one", "two", "three"}
	arr_value = []string{"a", "a", "a", "b", "b"}
)

func main() {

	for count := 0; count < 5; count++ {
		if arr_value[count] == "a" {
			fmt.Println(arr_key[count], "is", arr_value[count])
		} else {
			fmt.Println(arr_key[count], "??? that's not a!!")
		}
	}
}
