package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

var DB *gorm.DB

func main() {
	var err error
	dsn := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&User{})

	r := gin.Default()

	r.POST("/register", RegisterEndpoint)
	r.POST("/login", LoginEndpoint)

	// 非同期でサーバーを起動(Requestを送るため)
	go func() {
		r.Run()
	}()
	// サーバーが起動するまで待つ
	time.Sleep(1 * time.Second)

	// テストリクエスト
	TestRequest()
}

func RegisterEndpoint(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}

func LoginEndpoint(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var result User
	if err := DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&result).Error; err != nil {
		c.JSON(400, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful"})
}

func TestRequest() {
	// Resty clientを作成
	client := resty.New()

	// ユーザー登録
	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"username": "username",
			"password": "password",
		}).
		Post("http://localhost:8080/register")
	if err != nil {
		panic(err)
	}
	println("Response Info:")
	println("  Error      :", err)
	println("  Status     :", resp.Status())
	println("  Body       :\n", resp.String())
	println("--------------------")

	// ログイン
	resp, err = client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"username": "username",
			"password": "password",
		}).
		Post("http://localhost:8080/login")
	if err != nil {
		panic(err)
	}
	println("Response Info:")
	println("  Error      :", err)
	println("  Status     :", resp.Status())
	println("  Body       :\n", resp.String())
	println("--------------------")

	// 間違ったパスワードでログイン
	resp, err = client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"username": "username",
			"password": "wrong_password",
		}).
		Post("http://localhost:8080/login")
	if err != nil {
		panic(err)
	}
	println("Response Info:")
	println("  Error      :", err)
	println("  Status     :", resp.Status())
	println("  Body       :\n", resp.String())
	println("--------------------")
}
