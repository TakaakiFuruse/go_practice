// ####### hash in hash ########
package main

import "fmt"

func main() {
	m := map[string]map[string]int{
		"a": map[string]int{
			"1": 13,
			"2": 27,
		},
		"b": map[string]int{
			"42":  42,
			"109": 333,
		},
	}
	fmt.Printf("m[%s][%s] = %d\n", "a", "1", m["a"]["1"])
	fmt.Printf("m[%s][%s] = %d\n", "b", "42", m["b"]["42"])
}

//                ########Hash############
// package main

// import (
// 	"fmt"
// )

// type NinjaData struct {
// 	soul, style string
// }

// func main() {
// 	ninja := map[string]NinjaData{
// 		"NinjaSlayer":     {"Naraku", "Karate"},
// 		"Suicide":         {"Punk Ninja", "Psykic"},
// 		"Forest Sawatari": {"Vetcon Ninja", "Karate"},
// 	}

// 	fmt.Println(ninja)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	ar := [6]int{1, 2, 3, 4, 5, 6}
// 	fmt.Println(ar[0])
// 	fmt.Println(ar[:3])
// 	fmt.Println(ar[2:4])

// }

//              ############ Array #############
// package main

// import (
// 	"fmt"
// )

// const (
// 	Test1 = 1
// 	Test2 = 2
// 	Test3 = 3
// 	Hl    = "Hello!"
// )

// var (
// 	const_arr = []int{Test1, Test2, Test3}
// )

// func main() {
// 	for i := 0; i < len(const_arr); i++ {
// 		fmt.Println(Hl)

// 		fmt.Println(const_arr[i], "----", Hl)
// 	}

// }
