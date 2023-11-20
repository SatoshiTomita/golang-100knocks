package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addAll(arg ...int) int {
	sum := 0
	for _, n := range arg {
		sum += n
	}
	return sum
}

func main() {
	fmt.Println(add(1, 2)) // 3

	fmt.Println(addAll(1, 2, 3, 4, 5)) // 15
}
