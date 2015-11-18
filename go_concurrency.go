// package main

// import (
// 	"fmt"
// 	"time"
// )

// func threeSecond() {
// 	time.Sleep(3000 * time.Millisecond)
// 	fmt.Println("3 seconds passed")
// }

// func oneSecond() {
// 	time.Sleep(1000 * time.Millisecond)
// 	fmt.Println("1 second passed")
// }

// func twoSecond() {
// 	time.Sleep(2000 * time.Millisecond)
// 	fmt.Println("2 second passed")
// }

// func main() {

// 	fmt.Println(time.Now())
// 	go threeSecond()
// 	oneSecond()
// 	twoSecond()
// 	fmt.Println(time.Now())
// }

//-------------------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

func threeSecond() {
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("3 seconds passed")
}

func oneSecond() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("1 second passed")
}

func twoSecond() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("2 second passed")
}

func main() {
	funcArr := []func(){
		func() {
			threeSecond()
		},
		func() {
			twoSecond()
		},
		func() {
			oneSecond()
		},
	}

	var waitGroup sync.WaitGroup

	fmt.Println(time.Now())
	for _, value := range funcArr {
		waitGroup.Add(1)
		go func(function func()) {
			defer waitGroup.Done()
			function()
		}(value)
	}
	waitGroup.Wait()
	fmt.Println(time.Now())
}

//-------------------------------------------

package main

import (
  "fmt"
  "log"
  "time"
)

func wait1(c chan string) {
  time.Sleep(1 * time.Second)
  log.Print("waited 1 sec")
  c <- "wait 1 finished"

}

func wait2(c chan string) {
  time.Sleep(2 * time.Second)
  log.Print("waited 2 sec")
  c <- "wait 2 finished"
}

func wait3(c chan string) {
  time.Sleep(3 * time.Second)
  log.Print("waited 3 sec")
  c <- "wait 3 finished"
}

func main() {
  c := make(chan string)
  log.Print("started")
  go wait1(c)
  go wait2(c)
  go wait3(c)
  w1, w2, w3 := <-c, <-c, <-c
  log.Print("finished")
  fmt.Println(w1, w2, w3)
}

