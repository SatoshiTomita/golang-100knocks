package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "secret_key"

func main() {
	r := gin.Default()

	// ダミーのHTTPサーバーを起動
	api := r.Group("/api")
	{
		restricted := api.Group("/restricted")
		restricted.Use(AuthMiddleware) // 認証が必要なエンドポイントにミドルウェアを設定
		{
			restricted.GET("/profile", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "profile",
				})
			})

		}
	}

	// 非同期でサーバーを起動(Requestを送るため)
	go func() {
		r.Run()
	}()
	// サーバーが起動するまで待つ
	time.Sleep(1 * time.Second)

	// Resty clientを作成
	client := resty.New()

	// JWT トークンを使わずにリクエスト
	resp, err := client.R().
		EnableTrace().
		Get("http://localhost:8080/api/restricted/profile")
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Body       :\n", resp)
	fmt.Println("--------------------")

	// JWT トークンを使ってリクエスト
	tokenString := CreateToken(SECRET_KEY)
	resp, err = client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+tokenString).
		Get("http://localhost:8080/api/restricted/profile")
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Body       :\n", resp)
	fmt.Println("--------------------")

	// 間違ったJWT トークンを使ってリクエスト
	wrongTokenString := CreateToken("wrong_secret_key")
	resp, err = client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+wrongTokenString).
		Get("http://localhost:8080/api/restricted/profile")
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Body       :\n", resp)
	fmt.Println("--------------------")

}

// JWT トークンの生成
func CreateToken(secretKey string) string {
	encoded_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := encoded_token.SignedString([]byte(secretKey))
	fmt.Println("tokenString:", tokenString, "err:", err)
	return tokenString
}

// 認証が必要なエンドポイントのミドルウェア
func AuthMiddleware(c *gin.Context) {
	// ヘッダーからトークンを取得
	authorizationHeader := c.Request.Header.Get("Authorization")
	log.Println("authorizationHeader:", authorizationHeader)

	// トークンがない場合はエラー
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization header is required and must start with 'Bearer '",
		})
		c.Abort()
		return
	}
	tokenString := authorizationHeader[len("Bearer "):]
	fmt.Println("tokenString:", tokenString)

	// tokenの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if token.Valid {
		fmt.Println("正しいトークンです")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("これはトークンではありません")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("有効期限切れのトークンです")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
	} else {
		fmt.Println("トークンが無効です:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
	}

	c.Next()
}
