package main

import "fmt"

func main() {
	var n string
	var ch string
	// var count1 int = 0
	// var count2 int = 0
	// var m string
	fmt.Println("enter the expression:")
	fmt.Scan(&n)
	// length := len(n)

	stack := new(Stack)
	fmt.Println("matches and mismatches:")

	for i := 0; i < len(n); i++ {
		ch = string(n[i])

		if ch == "(" {

			stack.Push(ch)
			// count1++
			//fmt.Println(stack.Len())
		} else if ch == ")" {
			m := stack.Pop()
			// count2++
			fmt.Println(m)
			// fmt.Println("')' at index  %v matched with '(' at index ", ch, m)

		}

	}
	//fmt.Println(count1)
	//fmt.Println(count2)
	// if count1 == count2 {
	// 	fmt.Println("valid expresion")
	// } else {
	// 	fmt.Println("invalid expresion")
	// }
	if stack.Len() == 0 {
		fmt.Println("valid expresion")
	} else {
		fmt.Println("invalid expresion")
	}

}
