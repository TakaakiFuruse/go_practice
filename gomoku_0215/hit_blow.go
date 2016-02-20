package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	randAr := []string{}
	ansAr := []string{}
	hit := 0
	blow := 0

	for i := 1; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(nums))
		randAr = append(randAr, nums[n])
		nums = append(nums[:n], nums[n+1:]...)
	}

	for hit < 4 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		ansAr = strings.Split(text, "")

		for i := 0; i < 4; i++ {
			if randAr[i] == ansAr[i] {
				hit++
			}
			for j := 0; j < 4; j++ {
				if randAr[i] == ansAr[j] && randAr[i] != ansAr[i] {
					blow++
				}
			}
		}

		fmt.Println("hit is", hit)
		fmt.Println("blow is", blow)
		if hit == 4 {
			fmt.Println("WELL DONE!!")
			break
		}
		hit = 0
		blow = 0
	}
}
