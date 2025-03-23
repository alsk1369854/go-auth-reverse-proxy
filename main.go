package main

import (
	"flag"
	"go-auth-reverse-proxy/internal/pkg/models"
	"go-auth-reverse-proxy/internal/pkg/utils"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

// ===== 可調整參數區 =====
var (
	port string
	proxyURL string
	authFile string
)

var (
	authMap map[string]models.AuthToken
	// secretKey []byte
)

func getIPAddress(r *http.Request) string {
	// 支援經過反向代理（例如 Nginx）
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0]) // 取第一個真實 IP
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // fallback
	}
	return ip
}

// 驗證 Token 的簡單邏輯（可改成從 DB 查詢或 JWT 驗證）
func isAuthorized(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	val, exist := authMap[token]
	if exist {
		log.Printf("[AUTH] Pass name: %s\n", val.Name)
	} else {
		ip := getIPAddress(r)
		destination := r.URL.String()
		log.Printf("[AUTH] Reject IP: %s, Target: %s, Token:%s\n", ip, destination, token)
	}
	return exist
}


func init() {
	// 參數編譯
	flag.StringVar(&port, "port", "80", "HTTP redirect port")
	flag.StringVar(&proxyURL, "proxy", "http://localhost:8080", "Proxy target URL")
	flag.StringVar(&authFile, "auth", "auth-tokens.json", "Auth tokens json file")
	flag.Parse()

	// 打印當前工作目錄
	cwd, err := os.Getwd()
	if err != nil{
		log.Fatalf("get work directory error: %v\n", err)
	}
	log.Println("Work Directory:", cwd)

	// // 載入 .env 檔案
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalln("Unable to load .env file")
	// }

	// // 讀取 jwt secret key
	// secret := os.Getenv("JWT_SECRET")
	// if secret == ""{
	// 	log.Panicln("Environment value JWT_SECRET is not set!")
	// }
	// secretKey = []byte(secret)

	// 讀取 auth-token 配置文件
	tokens, err := utils.ReadAuthTokenFile(authFile)
	if err != nil{
		log.Fatalln(err)
	}

	// 建立 token map
	authMap = make(map[string]models.AuthToken)
	for _, t := range tokens {
		authMap[t.Token] = t
	}
}


func main() {
	proxy := utils.CreateProxy(proxyURL)
	listenAddr := "0.0.0.0:" + port
	
	// ===== 啟動 HTTP =====
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !isAuthorized(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		log.Println("[HTTP] Proxying:", r.URL.Path)
		proxy.ServeHTTP(w, r)
	})

	log.Printf("HTTP Listen on http://%s\n", listenAddr)
	log.Printf("Proxy Target %s\n", proxyURL)
	err := http.ListenAndServe(listenAddr, httpMux)
	if err != nil {
		log.Fatalf("HTTP Listen Error: %v", err)
	}
}

