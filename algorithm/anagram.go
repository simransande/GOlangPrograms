package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	var n string
	var m string
	var x string
	var y string

	var l string
	var P string
	// var str1 string
	// var str string
	fmt.Print("Enter one string  : ")
	fmt.Scan(&n)
	// str = strings.Replace(n, "", "s", -1)

	fmt.Print("Enter second string : ")
	fmt.Scan(&m)
	// str1 = strings.Replace(m, "", "s", -1)

	// length := len(m)
	// fmt.Println(length)

	// length1 := len(n)
	// fmt.Println(length1)

	var a []string = strings.Split(n, "")
	sort.Strings(a)
	//fmt.Println("Strings:", a)
	x = strings.Join(a, "")
	//fmt.Println(x)
	P = strings.ToLower(x)

	var a1 []string = strings.Split(m, "")
	sort.Strings(a1)
	//fmt.Println("Strings:", a1)
	y = strings.Join(a1, "")
	//fmt.Println(y)
	l = strings.ToLower(y)

	if P == l {
		fmt.Println("strings are anagram")

	} else {
		fmt.Println("strings are not anagram")
	}

}
