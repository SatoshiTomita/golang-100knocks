package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Hello, %s!\n", name)
}
