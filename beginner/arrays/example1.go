package main

import (
	"fmt"
)

var users = [...]string{"Luis", "Jorge", "Oscar"}

func main() {
	fmt.Println(users)
	for _, value := range users {
		// fmt.Println(key)
		fmt.Println(value)
	}
}