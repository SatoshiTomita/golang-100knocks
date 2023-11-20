package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// PostgreSQLへの接続
	dsn := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// スキーマの作成
	db.AutoMigrate(&Product{})

	// レコードの作成
	db.Create(&Product{Code: "D42", Price: 100})

	// レコードの取得
	var product Product
	// db.First(&product, 1)                 // idが1のレコードを取得
	db.First(&product, "code = ?", "D42") // codeがD42のレコードを取得
	fmt.Println("Created:", product.ID, product.Code, product.Price)

	// レコードの更新
	db.Model(&product).Update("Price", 200)
	fmt.Println("Updated:", product.ID, product.Code, product.Price)

	// 複数のレコードの更新
	db.Model(&product).Updates(Product{Price: 300, Code: "F42"}) // 構造体で更新
	fmt.Println("Updated:", product.ID, product.Code, product.Price)
	db.Model(&product).Updates(map[string]interface{}{"Price": 400, "Code": "G42"})
	fmt.Println("Updated:", product.ID, product.Code, product.Price)

	// レコードを削除
	db.Delete(&product, "code = ?", "G42")
}
