package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1から10までの合計:", sum)

	sum = 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println("2のべき乗:", sum)

	list := []string{"a", "b", "c"}
	for index, item := range list {
		fmt.Println(index, item)
	}
}
