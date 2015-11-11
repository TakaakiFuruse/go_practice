package main

import (
	"fmt"
	"sort"
)

// fizbuz Map  -> pair of num and string
// fizbuz arra => fill by hash key -> how to put key into array??
// if num / (fizubuz array num) -> return fizbuz[fizubuz array num] string

var (
	fizBuzMap = map[int]string{
		1:  "1",
		3:  "multiple of 3",
		5:  "multiple of 5",
		15: "multiple of 15",
	}
	quotientArray = []int{15, 5, 3}
)

func returnString(num int) string {
	sort.Sort(sort.Reverse(sort.IntSlice(quotientArray)))
	v := ""
	for _, value := range quotientArray {
		// fmt.Println("value:", value, ",", "num:", num, ",", "num%value", num%value)
		if value == 1 || num%value == 0 {
			v = fizBuzMap[value]
			// fmt.Println("hash is", fizBuzMap[value])
		} else {
			v = fmt.Sprint(num)
		}
	}
	return v
}

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Println(returnString(i))
	}

}
