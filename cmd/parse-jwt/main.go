package main

import (
	"flag"
	"fmt"
	"go-auth-reverse-proxy/internal/pkg/utils"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	secretKey []byte
	token string
)

func init(){
	flag.StringVar(&token, "token", "", "Parse Token")
	flag.Parse()

	// 載入 .env 檔案
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env file")
	}

	secretKey = []byte(os.Getenv("JWT_SECRET"))
}


func main(){
	claims, err := utils.ParseJWT(token, secretKey)
	if err != nil{
		log.Fatalf("Parse Error %v", err)
	}

	fmt.Println("Decode content:")
	for key, val := range claims {
		fmt.Printf("  %s: %v\n", key, val)
	}
	fmt.Printf("Issued at: %s", time.Unix(int64(claims["iat"].(float64)), 0))
}