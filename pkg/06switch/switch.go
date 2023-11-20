package main

import "fmt"

func main() {
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("月曜日")
	case "Tuesday":
		fmt.Println("火曜日")
	case "Wednesday":
		fmt.Println("水曜日")
	default:
		fmt.Println("その他")
	}
}
