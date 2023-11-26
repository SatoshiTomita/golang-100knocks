package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(hmacSampleSecret []byte) string {
	// JWT トークンの生成
	encoded_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := encoded_token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)
	return tokenString
}

func VerifyToken(tokenString string, hmacSampleSecret []byte) bool {
	// トークンの検証
	decoded, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // 署名アルゴリズムの検証
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil // シークレットキーの検証
	})
	if err != nil {
		log.Fatal(err)
		return false
	}

	// トークンの中身を取得
	if claims, ok := decoded.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["foo"], claims["nbf"])
		return true
	} else {
		fmt.Println(err)
		return false
	}
}

func main() {
	hmacSampleSecret := []byte("weak_key") // シークレットキー

	// JWT トークンの生成
	tokenString := CreateToken(hmacSampleSecret)

	// トークンの検証
	VerifyToken(tokenString, hmacSampleSecret)
}
