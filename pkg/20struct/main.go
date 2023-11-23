package main

import "fmt"

// Personは人物の情報を表す構造体です。
type Person struct {
	Name    string
	Age     int
	Address Address
}

// Addressは住所を表す構造体です。
type Address struct {
	City    string
	Street  string
	ZipCode string
}

func main() {
	// 構造体の初期化
	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:    "Tokyo",
			Street:  "123 Main St",
			ZipCode: "100-0001",
		},
	}

	// 構造体のフィールドにアクセス
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("City: %s\n", p.Address.City)
	fmt.Printf("Street: %s\n", p.Address.Street)
	fmt.Printf("Zip Code: %s\n", p.Address.ZipCode)
}
