package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	randAr := []int{}

	for i := 1; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(nums))
		randAr = append(randAr, nums[n])
	}

	fmt.Println(randAr)

}
