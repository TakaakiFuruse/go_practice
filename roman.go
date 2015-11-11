package main

import (
	"fmt"
	s "strings"
)

var (
	Romans_ones = map[int]string{
		1: "i",
		2: "nil",
		3: "nil",
		4: "iv",
		5: "v",
		6: "nil",
		7: "nil",
		8: "nil",
		9: "ix",
	}

	Romans_tens = map[int]string{
		1: "x",
		2: "nil",
		3: "nil",
		4: "xl",
		5: "l",
		6: "nil",
		7: "nil",
		8: "nil",
		9: "xc",
	}

	Romans_hundreads = map[int]string{
		1: "c",
		2: "nil",
		3: "nil",
		4: "cd",
		5: "d",
		6: "nil",
		7: "nil",
		8: "nil",
		9: "cm",
	}
	Arr_index_to_Keta_num = map[int]map[int]string{
		0: Romans_hundreads,
		1: Romans_tens,
		2: Romans_ones,
	}
)

// func iterateSplit(array, []string{}) {

// }

func main() {
	Split := s.Split("332", "")
	for key, _ := range Split {
		fmt.Println(Arr_index_to_Keta_num[key])
	}
}
