package main

import (
	"flag"
	"fmt"
	"go-auth-reverse-proxy/internal/pkg/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	name string
	secretKey string
)

func init() {
	flag.StringVar(&name, "name", "admin", "Name of Token")
	flag.Parse()

	// 載入 .env 檔案
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env file")
	}

	secretKey = os.Getenv("JWT_SECRET")
}

func main(){
	token, err  := utils.GenerateJWT(name, secretKey); if err != nil{
		log.Fatalf("GenerateJWT Error %v\n", err)
	}
	fmt.Println(token)
}