package utils

import (
	"encoding/json"
	"fmt"
	"go-auth-reverse-proxy/internal/pkg/models"
	"os"
)


func ReadAuthTokenFile(path string) ([]models.AuthToken, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable read file %s: %w", path,  err)
	}
	
	// 解析 JSON 為 slice
	var tokens []models.AuthToken
	if err = json.Unmarshal(data, &tokens); err != nil {
		return nil, fmt.Errorf("json parsing error: %w", err)
	}

	return tokens, nil

	// // 建立 map：token 為 key
	// tokenMap := make(map[string]AuthToken)
	// for _, t := range tokens {
	// 	tokenMap[t.Token] = t
	// }

	// // 測試印出
	// for token, info := range tokenMap {
	// 	fmt.Printf("Token: %s\nUser: %s\n\n", token, info.Username)
	// 	}
}


