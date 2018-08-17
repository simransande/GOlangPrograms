package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// type Users struct {
// 	Users []Users}
func main() {

	str := "Hello <<name>>, We have your full name as <<full>> <<Name>> in our system. your contact number is 91-xxxxxxxxxx. Please,let us know in case of any clarification Thank you BridgeLabz 2016-01-01."
	fmt.Println(str)
	var name string
	fmt.Println("Enter your name:")
	fmt.Scan(&name)

	var lname string
	fmt.Println("Enter your last name:")
	fmt.Scan(&lname)

	var num int
	fmt.Println("enter your mobile number:")
	fmt.Scan(&num)
	t := strconv.Itoa(num)
	//fmt.Println(t)

	currentTime := time.Now()
	m := currentTime.Format("2006-01-02")

	str = strings.Replace(str, "<<name>>", name, -1)
	str = strings.Replace(str, "<<full>>", name, -1)
	str = strings.Replace(str, "<<Name>>", lname, -1)
	str = strings.Replace(str, "xxxxxxxxxx", t, -1)
	str = strings.Replace(str, "2016-01-01", m, -1)
	fmt.Println(str)

}
