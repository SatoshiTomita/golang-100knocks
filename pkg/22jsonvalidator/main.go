package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Username string `validate:"required,printascii,max=16"`
	Name     string
	Email    string `validate:"required,email"`
	Age      int    `validate:"min=0"`
}

func main() {
	// 失敗する構造体の初期化
	alice := User{
		Username: "alice",
		Name:     "Alice",
		Email:    "",
		Age:      -1,
	}
	result := validateUser(alice)
	fmt.Println(result)

	// 成功する構造体の初期化
	bob := User{
		Username: "bob",
		Name:     "Bob",
		Email:    "sample@email.com",
		Age:      25,
	}

	// バリデーション
	result = validateUser(bob)
	fmt.Println(result)
}

// バリデーション
func validateUser(user User) string {
	fmt.Println(user)

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return err.Error()
	} else {
		return "ok"
	}
}
