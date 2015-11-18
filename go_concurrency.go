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
