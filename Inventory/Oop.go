package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"Inventory"`
}

type User struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Price  int    `json:"price"`
}

func main() {

	jsonFile, err := os.Open("Inventory.json") // Open jsonFile

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Inventory.json")
	// defer jsonFile.Close()                   //defer- file closed that we can parse
	byteValue, _ := ioutil.ReadAll(jsonFile) //read our opened file as byte array

	// we initialize our Users array
	var users Users

	json.Unmarshal(byteValue, &users) //encode the data

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("Name: " + users.Users[i].Name)
		fmt.Println("Weight: " + strconv.Itoa(users.Users[i].Weight))
		fmt.Println("Price: " + strconv.Itoa(users.Users[i].Price))
	}

}
