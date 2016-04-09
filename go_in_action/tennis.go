package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)
	wg.Add(2)

	go player("A", court)
	go player("B", court)

	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("%s won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s Missed\n", name)
			close(court)
			return
		}

		fmt.Printf("player %s hit %d\n", name, ball)
		ball++

		court <- ball
	}
}
