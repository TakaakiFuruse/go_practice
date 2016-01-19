package main

import (
	"fmt"
	"log"
	"time"
)

func wait1(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "1"
}

func wait2(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "2"
}

func wait3(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "3"
}

func main() {
	c := make(chan string)
	log.Print("started")
	go wait1(c)
	go wait2(c)
	go wait3(c)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c:
			fmt.Println("finished", msg1)
		case msg2 := <-c:
			fmt.Println("finished", msg2)
		case msg3 := <-c:
			fmt.Println("finished", msg3)

		}
	}

	log.Print("finished")

}
