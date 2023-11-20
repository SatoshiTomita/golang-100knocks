package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQLドライバのインポート
)

type User struct {
	ID      int64
	Name    string
	Profile sql.NullString
	Created time.Time
	Updated time.Time
}

func main() {
	// PostgreSQLへの接続
	dbDriver := "postgres"
	dsn := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := setupDB(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// スキーマの作成
	err = createUsersTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// データの挿入
	user := User{
		Name: "John Doe",
	}
	err = insertUser(db, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted user:", user)

	// データの読み取り
	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users:", users)

	// データの更新
	user.Name = "John Doe Updated"
	err = updateUser(db, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated user:", user)

	// データの削除
	err = deleteAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all users")
}

func setupDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			profile TEXT,
			created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	return nil
}

func insertUser(db *sql.DB, user *User) error {
	sqlStatement := "INSERT INTO users (name) VALUES ($1) RETURNING id"
	err := db.QueryRow(sqlStatement, user.Name).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func updateUser(db *sql.DB, user *User) error {
	sqlStatement := "UPDATE users SET name = $2 WHERE id = $1"
	_, err := db.Exec(sqlStatement, user.ID, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func deleteAllUsers(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		return err
	}
	return nil
}
