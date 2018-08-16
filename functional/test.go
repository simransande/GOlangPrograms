package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mains() {

	var headcount int = 0
	var tailcount int = 0
	var per1 int = 0
	var per2 int = 0
	var n int
	//var sum float32 = 0
	fmt.Print("Enter a number how much time you want to flip: ")
	fmt.Print("Enter a number how much time you want to flip: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ { // assigning 1 to i
		myrand := randIntt(1, 6)
		fmt.Println(myrand)
		if myrand > 3 {
			headcount++
		} else {
			tailcount++
		}
	}

	fmt.Println("number of heads are %d", headcount)
	fmt.Println("number of tails are %d", tailcount)

	per1 = headcount / n
	per2 = tailcount / n

	fmt.Println("percentage of headcount is %v", per1)
	fmt.Println("percentage of tailcount is %v", per2)
}

func randIntt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
