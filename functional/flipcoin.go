package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var headcount int = 0
	var tailcount int = 0
	var per1 int = 0
	var per2 int = 0
	var n int
	//var sum float32 = 0
	fmt.Print("Enter a number how much time you want to flip: ")
	//fmt.Print("Enter a number how much time you want to flip: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ { // assigning 1 to i
		myrand := randInt(1, 6)
		fmt.Println(myrand)
		if myrand > 3 {
			headcount++
		} else {
			tailcount++
		}
	}

	fmt.Println("num", headcount)
	fmt.Println("number of tails are %d", tailcount)
	//fmt.Println(headcount)
	//fmt.Println(tailcount)
	//fmt.Println(n)

	per1 = headcount * 100 / n
	//fmt.Println(per1)
	per2 = tailcount * 100 / n

	fmt.Println("percentage of headcount is %v", per1)
	fmt.Println("percentage of tailcount is %v", per2)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
