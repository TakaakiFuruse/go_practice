package main

import "fmt"

func fizbuz(x int) int {
	if divisible(x, 15) {
		fizbuz15()
	}
	if divisible(x, 5) {
		fizbuz5()
	}
	if divisible(x, 3) {
		fizbuz3()
	} else {
		fmt.Println(x)
	}
	return x
}

func divisible(num1, num2 int) bool {
	return num1%num2 == 0
}

func fizbuz15() {
	fmt.Println("fizbuz")
}

func fizbuz5() {
	fmt.Println("fiz")
}

func fizbuz3() {
	fmt.Println("buz")
}

func main() {
	for i := 1; i < 16; i++ {
		fizbuz(i)
	}
}
