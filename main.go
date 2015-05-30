package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]
	if command == "login" {
		username := os.Args[2]
		password := os.Args[3]
		result, err := Login(username, password)
		fmt.Println(result)
		if err != nil {
			fmt.Println(err)
		}
	} else if command == "logout" {
		result, err := Logout()
		fmt.Println(result)
		if err != nil {
			fmt.Println(err)
		}
	}
}
