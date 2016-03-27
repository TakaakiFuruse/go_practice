package main

import "fmt"

var (
	instance_var string = "i am kinda instance var"
)

func print_nums() {
	var a int = 40
	var b int = 30
	var c int = 40
	fmt.Println("a and b  is", a == b)
	fmt.Println("a == b || a == c  is", a == b || a == c)
	fmt.Println("a && b || a && c  is", a == b && a == c)
	fmt.Println("not true ", !true)
	fmt.Println("a c  not true ", a != c)
	fmt.Println("p instance var from print_nums", instance_var)

	instance_var = "i am from print nums"
	fmt.Println("insatnce_var from print_nums", instance_var)

}

func for_and_if_statement() {
	for i := 0; i <= 10; i += 1 {
		if i == 5 {
			fmt.Println("It's 5 !! Congrats!!")
		} else if i == 8 {
			fmt.Println("it's 8!!!")
		} else {
			fmt.Println("it's something, whatever..")
		}

	}

}

func forAndSwith() {

	fmt.Println("instance_var from forAndSwitch", instance_var)
	for i := 0; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println("this is 1")
		case 5:
			fmt.Println("this is 5")
		default:
			fmt.Println(i)
		}
	}
}

// =====================================================================

func myArray() {

	var Array2 [5]float64
	Array2[0] = 6
	Array2[1] = 7
	Array2[2] = 8
	Array2[3] = 9
	Array2[4] = 10
	Array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array2", Array2[3])
	fmt.Println("array", Array[3])

	for _, value := range Array2 {
		fmt.Println(value)
	}
}

func mySlice() {
	numSlice := []int{5, 4, 3, 2, 1}
	//numSlice2 := numSlice[0:3] //5,4,3
	// numSlice2 := numSlice[3:] //2,1
	numSlice2 := numSlice[:3] //5,4,3
	fmt.Println("numSlice2[0] =", numSlice2)
}

func mySliceMake() {
	numSlice := []int{1, 2, 3, 4, 5, 6}
	numSlice3 := make([]int, 6, 10)
	copy(numSlice3, numSlice)
	numSlice3 = append(numSlice3, 7, 8, 9, 10, 11, 12)
	fmt.Println(numSlice3)
}

func myMap() {
	age := make(map[string]int)
	age["Masashi Tashiro"] = 50
	fmt.Println(age)
	age["Your Highness Demon Kogure"] = 100056
	fmt.Println(age)
	age["AAA"] = 111
	fmt.Println(age)
	delete(age, "Your Highness Demon Kogure")
	fmt.Println(age)
}

// =====================================================================

func returnValue1() {
	numArray := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println("sum of numarray is ", addUp(numArray))
}

func addUp(numbersArray []int) int {
	sum := 0
	for _, value := range numbersArray {
		sum += value
	}
	return sum
}

func returnTwoValues() {

	var1, var2 := addValue("I am Var!")
	fmt.Println(var1)
	fmt.Println(var2)
}

func addValue(number string) (string, string) {
	return number + "I am the first.", number + "I am the second."
}

func splatOperater(args ...string) string {
	resultVar := ""
	for _, value := range args {
		resultVar += value + " "
	}
	return resultVar
}

// =====================================================================

func myClosuer() {
	string1 := "Hello"
	string2 := " Monday!!"

	// func addString(){} does not work
	addString := func() string {
		string1 += string2
		string1 += " and Goodbye Sunday"
		return addOrz(string1)
	}
	fmt.Println(addString())
}

func addOrz(arg string) string {
	arg += " ... orz"
	return arg
}

// =====================================================================

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// =====================================================================

func printOne() { fmt.Println("i'm the first!!") }
func printTwo() { fmt.Println("i'm the second!!") }

func printOneAndTwo() {
	defer printTwo() // -> without defer, Two comes first
	printOne()
}

// =====================================================================

func returnString(arg1, arg2 int) int {

	// if there were no recover go would "panic".
	defer func() {
		fmt.Println(recover())
	}()
	result := arg1 / arg2
	return result
}

// =====================================================================

func thisisNtPointer(arg string) {
	arg = "I love Ruby, not GO!!"
}

func thisisPointer(arg *string) {
	*arg = "I love GO!!"
}

// =====================================================================

type CalorieCalc struct {
	daiJiro              int
	largeButaDoubleYasai int
	largeButaYasai       int
	normal               int
	mini                 int
	coke                 int
	woolongTea           int
}

func (reciever *CalorieCalc) sumoWrestler() int {
	return reciever.daiJiro + reciever.coke
}

func (reciever *CalorieCalc) aGuy() int {
	return reciever.largeButaYasai + reciever.woolongTea
}

// ===================================================
func main() {
	// print_nums()
	// forAndSwith()
	// for_and_if_statement()

	// myArray()
	// mySlice()
	// mySliceMake()
	// myMap()

	// returnValue1()
	// returnTwoValues()
	// fmt.Println(splatOperater("i", "LUV", "ruby",
	//  "and", "soso", "love", "go !!"))
	// fmt.Println(factorial(10))

	// defer printTwo() // -> without defer, Two comes first
	// printOneAndTwo()

	// fmt.Println(returnString(4, 2))
	// fmt.Println(returnString(2, 0))

	// argStr := "I love ... what?"
	// fmt.Println(argStr)
	// thisisNtPointer(argStr)
	// fmt.Println(argStr)
	// thisisPointer(&argStr) // need "&" here!!
	// fmt.Println(argStr)

	// calories := CalorieCalc{3000, 2000, 1800, 1000, 800, 500, 0}
	// fmt.Println("You took", calories.daiJiro, "kilo calories")
	// fmt.Println("You took", calories.mini, "kilo calories")
	// fmt.Println("SumoWrestler ate", calories.sumoWrestler(), "kilo calories")
	// fmt.Println("A GUy ate", calories.aGuy(), "kilo calories")

}
