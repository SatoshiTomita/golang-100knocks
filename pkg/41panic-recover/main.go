package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()

	fmt.Println("This is a statement before panic.")
	panic("This is a panic situation!")
	// fmt.Println("This statement will not be executed.")
}
