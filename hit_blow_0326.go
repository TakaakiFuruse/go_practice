package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Game defines game
type Game struct {
	randAr []string
	ansAr  []string
	hit    int
	blow   int
}

func (g *Game) randArGenerator() []string {
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 1; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(len(nums) - 1)
		g.randAr = append(g.randAr, nums[randNum])
		nums = append(nums[:randNum], nums[randNum+1:]...)
	}
	return g.randAr
}

func (g *Game) ansGetter() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	g.ansAr = strings.Split(text, "")
	return g.ansAr
}

func (g *Game) ansChecker() {
	g.hit = 0
	g.blow = 0

	for i := 0; i < 4; i++ {
		if g.ansAr[i] == g.randAr[i] {
			g.hit++
		}
		for j := 0; j < 4; j++ {
			if g.ansAr[i] == g.randAr[j] && g.ansAr[i] != g.randAr[i] {
				g.blow++
			}
		}
	}
	if g.hit == 4 {
		fmt.Println("WELL DONE!!")
	} else {
		fmt.Printf("Hit is %v Blow is %v \n", g.hit, g.blow)
	}
}

func main() {
	var g Game
	g.randArGenerator()
	for g.hit < 4 {
		fmt.Println(g.randAr)
		g.ansGetter()
		g.ansChecker()
	}
}
