package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]
	if command == "login" {
		username := USERNAME
		password := PASSWORD
		if username == "" || password == "" {
			username = os.Args[2]
			password = os.Args[3]
		}
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
	} else if command == "status" {
		result, err := GetStatus()
		fmt.Println(result)
		if err != nil {
			fmt.Println(err)
		}
	} else if command == "keep" {
		username := os.Args[2]
		password := os.Args[3]
		err := Keep(username, password)
		if err != nil {
			fmt.Println(err)
		}
	}
}
