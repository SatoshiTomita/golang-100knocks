package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// PostgreSQLへの接続情報
	dsn := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// トランザクション内での処理（例: データの挿入）
	_, err = tx.Exec(`INSERT INTO mytable (column1, column2) VALUES ($1, $2)`, "value1", "value2")
	if err != nil {
		// エラーが発生した場合はロールバック
		tx.Rollback()
		log.Fatal(err)
	}

	// トランザクションのコミット
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction completed successfully!")
}
