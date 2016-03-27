package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
}

func myThread1(s string) {
	words := [3]string{"i am in the first",
		"i am in the second", "i am in the third"}
	for _, value := range words {
		say(value)
		go say(s)
	}
}

func myThread2(s string) {
	words := [3]string{"i am in the first",
		"i am in the second", "i am in the third"}
	for _, value := range words {
		go say(s)
		say(value)
	}
}

func myThread3(s string) {
	words := [3]string{"i am in the first",
		"i am in the second", "i am in the third"}
	for _, value := range words {
		go say(s)
		say(value)
	}
}

func main() {
	// myThread1("YOLO!")
	// myThread2("YOLO!")
	myThread3("YOLO!")
}
