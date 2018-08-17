package main

import "fmt"
import "io/ioutil"
import "strings"

func main() {
	b, err := ioutil.ReadFile("file.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	l := List{}

	//fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	//fmt.Println(str) // print the content as a 'string'
	stringSlice := strings.Split(str, " ")
	//fmt.Printf("%v\n", stringSlice)

	for index := 0; index < len(stringSlice); index++ {
		l.Insert(stringSlice[index])
	}

	l.Display()
	var n string
	fmt.Println("enter the name you wanst to search")
	fmt.Scan(&n)
	flag := l.Search(n)

	if flag {
		fmt.Println("Found")
		l.Delete(n)
		l.Display()
	} else {
		fmt.Println("Not found")
		l.Insert(n)
		l.Display()
	}

}
