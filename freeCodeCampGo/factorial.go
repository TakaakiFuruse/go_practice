package main

import "fmt"

func factorialize(t int) int {
	ar := []int{}
	for i := 1; i < t+1; i++ {
		ar = append(ar, i)
	}
	f := 1
	for _, value := range ar {
		f *= value
	}
	return f
}

func main() {
	fmt.Println(factorialize(5) == 120)

	fmt.Println(factorialize(10) == 3628800)

	fmt.Println(factorialize(20) == 2432902008176640000)

	fmt.Println(factorialize(0) == 1)

}
