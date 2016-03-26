package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Game struct {
	randAr []string
	ansAr  []string
	hit    int
	blow   int
}

func (ar *Game) randArGenerator() []string {
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 1; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(len(nums) - 1)
		ar.randAr = append(ar.randAr, nums[randNum])
		nums = append(nums[:randNum], nums[randNum+1:]...)
	}
	return ar.randAr
}

func (ar *Game) ansGetter() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	ar.ansAr = strings.Split(text, "")
	return ar.ansAr
}

func (g *Game) ansChecker() bool {
	g.hit = 0
	g.blow = 0
	for i := 0; i < 4; i++ {
		if g.ansAr[i] == g.randAr[i] {
			g.hit += 1
		}
		for j := 0; j < 4; j++ {
			if g.ansAr[i] == g.randAr[j] && g.ansAr[i] != g.randAr[i] {
				g.blow += 1
			}
		}
	}
	if g.hit == 4 {
		fmt.Println("WELL DONE!!")
	} else {
		fmt.Printf("Hit is %v Blow is %v \n", g.hit, g.blow)
	}
	return true
}

func main() {
	var g Game
	g.randArGenerator()
	for g.blow < 4 {
		fmt.Println(g.randAr)
		g.ansGetter()
		g.ansChecker()
	}
}
