package main

import (
	"fmt"
	"strconv"
	s "strings"
)

var (
	Romans_ones = map[string]string{
		"1": "i", "2": "nil", "3": "nil",
		"4": "iv",
		"5": "v", "6": "nil", "7": "nil", "8": "nil",
		"9": "ix",
	}

	Romans_tens = map[string]string{
		"1": "x", "2": "nil", "3": "nil",
		"4": "xl",
		"5": "l", "6": "nil", "7": "nil", "8": "nil",
		"9": "xc",
	}

	Romans_hundreads = map[string]string{
		"1": "c", "2": "nil", "3": "nil",
		"4": "cd",
		"5": "d", "6": "nil", "7": "nil", "8": "nil",
		"9": "cm",
	}
	Arr_index_to_Keta_num = map[string]map[string]string{
		"0": Romans_hundreads,
		"1": Romans_tens,
		"2": Romans_ones,
	}
)

func split_nums(number int) {
	s.Split(strconv.Itoa(number), "")
}

func main() {
	fmt.Println(Arr_index_to_Keta_num["0"]["4"])
	fmt.Println(split_nums(332))
}
