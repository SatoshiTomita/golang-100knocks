package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "yourpassword"

	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Println("ハッシュ化されたパスワード: ", string(hashedPassword))

	// パスワードが一致するかチェック
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		panic(err)
	}
	fmt.Println("パスワードが一致しました")
}
