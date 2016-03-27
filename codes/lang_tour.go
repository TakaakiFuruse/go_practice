//  =============1---17=================

package main

import (
	"fmt"
	"strconv"
)

func add(x int, y int, z int) int {
	return x + y + z
}

func add2(x, y, z int) int {
	return x + y + z
}

func swap(x, y string) (string, string) {
	return y, x

}

func nakedReturnSplit(sum int) (x, y int) {
	x = sum + 2
	y = sum + 100
	return
}

func ifStatement(str string) string {
	if str == "YO" {
		str = "WHAZZAAAAAP!"
	}
	return str
}

func ifStatement2(tes1, tes2, tes3 string) string {
	if tes1 == "hey" && tes2 == "yo" {
		return tes3
	}
	return tes3
}

func ifStatement3(tes1, tes2 int) int {
	if tes1++; tes1 < 3 {
		tes2 = 13
	}
	return tes2
}

func ifStatement4(tes1 int) string {
	var tes3 string
	if v := tes1 % 2; v == 0 {
		tes3 = "even num"
		return tes3
	} else {
		tes3 = "odd num"
		return tes3
	}
}

func switchStatment(num int) string {
	var v string
	switch {
	case num%15 == 0:
		v = "multiple of 15"
	case num%2 == 0:
		v = "multiple of 2"
	case num%3 == 0:
		v = "multiple of 3"
	case num%5 == 0:
		v = "multiple of 5"
	default:
		v = strconv.Itoa(num)
	}
	return v
}

func deffer1() {
	i := 0
	fmt.Println(i)
	i++
	defer fmt.Println(i)
	return
}

func deffer2() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
		fmt.Print(i, " ")
	}
}

func main() {

	for i := 0; i < 31; i++ {
		fmt.Println(switchStatment(i))
	}

	// ======= fizz buzz===========
	// for i := 1; i < 31; i++ {
	// 	switch {
	// 	case i%15 == 0:
	// 		fmt.Println("multiple of 15")
	// 	case i%2 == 0:
	// 		fmt.Println("multiple of 2")
	// 	case i%3 == 0:
	// 		fmt.Println("multiple of 3")
	// 	case i%5 == 0:
	// 		fmt.Println("multiple of 5")

	// 	default:
	// 		fmt.Println(i)
	// 	}
	// }

	// fmt.Println(ifStatement4(2))
	// fmt.Println(ifStatement4(3))
	// fmt.Println(ifStatement4(4))

	// fmt.Println(ifStatement3(1, 7)) // -> 13
	// fmt.Println(ifStatement3(2, 7)) // -> 7

	// fmt.Println(
	// 	ifStatement2("hey", "yo", "WAZAAAAAP!"),
	// 	ifStatement2("good day sir", "morning gents", "??????"),
	// )

	// fmt.Println(ifStatement("HEY"))
	// fmt.Println(ifStatement("YO"))

	// sum := "test"
	// i := 1
	// for i <= 10 {
	//   fmt.Println(sum)
	//   i += 1
	// }

	// fmt.Println(sum)

	// sum := "test"
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(sum)
	// }

	// sum := 1
	// for sum < 10 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// sum := 0
	// for i := 0; i <= 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// const Fi = "Fi"
	// const Pi = "Pi"
	// Fi = "Hey"
	// Pi = "Pye"
	// fmt.Println(Fi, Pi)

	// var x, y int = 3, 4
	// var f = math.Sqrt(float64(x*x + y*y))
	// fmt.Println(f)

	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Println(i, f, b, s)
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// fmt.Println(nakedReturnSplit(1))

	// a, b := swap("hello", "world")
	//  fmt.Println(swap(a, b))
	//  fmt.Println(swap(b, a))

	//  fmt.Printf("Now you have %g problems.",

	//    math.Nextafter(2, 3))

	//  fmt.Println("print number", rand.Intn(10))

	// fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))

	//  fmt.Println(math.Pi)

	// fmt.Println(add(1, 2, 3))
	// fmt.Println(add2(1, 2, 3))
}
