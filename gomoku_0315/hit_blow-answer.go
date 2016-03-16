package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	question := generateQuestion()

	try := 0
	for {
		try++
		hit, blow := answer(question)
		if hit == 4 {
			fmt.Println("Congratulation！")
			break
		} else {
			fmt.Println(try, "回目: ", "hit:", hit, "blow:", blow)
		}
	}
}

func generateQuestion() string {
	var nums []string
	var question string
	nums = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for len(question) < 4 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Intn(10)
		ok := check(question, nums[num])
		if ok {
			question += nums[num]
		}
	}

	return question
}

func check(question string, numStr string) bool {
	for i := range question {
		if string(question[i]) == numStr {
			return false
		}
	}

	return true
}

func answer(question string) (int, int) {
	var answer string
	var hit, blow int
	fmt.Print("数字を入れてね！(answer): ")
	_, err := fmt.Scan(&answer)

	if err != nil {
		fmt.Println("error!!")
	}

	if len(answer) > 4 {
		fmt.Println(os.Stderr, "四桁にしてください")
	} else if len(answer) < 4 {
		fmt.Println(os.Stderr, "4桁にしてください")
	} else {
		for i := range question {
			for k := range answer {
				if question[i] != answer[k] {
					continue
				}
				if i == k {
					hit++
				} else {
					blow++
				}
			}
		}
	}
	return hit, blow
}
