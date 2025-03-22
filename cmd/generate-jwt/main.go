package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var (
	username string
	secretKey []byte
)

func init() {
	flag.StringVar(&username, "username", "admin", "Username of Token")
	flag.Parse()

	// 載入 .env 檔案
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env file")
	}

	secretKey = []byte(os.Getenv("JWT_SECRET"))
}

func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"iat":      time.Now().Unix(),                   // 發行時間
		// "exp":      time.Now().Add(time.Hour * 1).Unix(), // 過期時間
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func main(){
	token, err  := generateJWT(username); if err != nil{
		log.Fatalf("GenerateJWT Error %v\n", err)
	}
	fmt.Println(token)
}