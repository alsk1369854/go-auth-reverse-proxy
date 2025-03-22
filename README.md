# go-auth-reverse-proxy

go-auth-reverse-proxy 是一款使用 Go 語言實作的輕量級 具授權驗證機制的反向代理伺服器。
透過 Bearer Token（靜態 JSON 或 JWT）驗證機制，保護後端服務僅允許授權使用者存取。
適合用於開發、內網服務保護或作為 API Gateway 前置防線使用。

- 支援 Token 驗證（靜態或 JWT）

- 🔁 請求轉發至自訂 Proxy URL

- ⚙️ CLI 與環境變數雙重配置支援

- 🐳 支援 Docker 一鍵部署

- 📜 提供簡易 JWT 生成工具

## Deploy

### Docker

```bash
docker run -d \
--name my-go-auth-reverse-proxy \
-p 80:80 \
-e PROXY_URL=http://host.docker.internal:8080 \
--restart unless-stopped \
alsk1369854/go-auth-reverse-proxy:0.0.0
```

## Dev

### Run app

```bash
go run main.go -port=80 -proxy=http://localhost:8080 -auth=auth-tokens.json
```

### Docker image

#### Build

```bash
docker build -t go-auth-reverse-proxy:0.0.0 -f deployments/Dockerfile .
```

#### Push to DockerHub

```bash
export tag=0.0.0
export docker_account=alsk1369854

docker login

docker tag go-auth-reverse-proxy:${tag} ${docker_account}/go-auth-reverse-proxy:${tag}

docker push ${docker_account}/go-auth-reverse-proxy:${tag}
```

## Generate JWT

```bash
# Create Secret Key
echo JWT_SECRET=$(openssl rand -base64 32) > .env

# Generate token
go run cmd/generate-jwt/main.go -username=your_username
```

## Test Proxy

### Create Test Server Container

```bash
docker run -d \
--name py-test-server \
-p 8080:80 \
--restart unless-stopped \
alsk1369854/py-test-server
```

### Test API list

```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NDI2MTQ1NDAsInVzZXJuYW1lIjoic2RwbWxhYiJ9.eIj96Wpa3yYVK_CDIOk3CM8K8EoEQEpdiF0YKu_TQac" http://localhost/
# Respond {"message":"Hello world"}

curl -H "Authorization: Bearer yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NDI2MTQ1NjIsInVzZXJuYW1lIjoidGVtcDEifQ.ZBIzEGKlry5kug4ZVx_KZSJfwHU8YaIRElhNRfxnKAo" http://localhost/user
# Respond {"message":"User"}

curl -H "Authorization: Bearer yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NDI2MTQ1NjIsInVzZXJuYW1lIjoidGVtcDEifQ.ZBIzEGKlry5kug4ZVx_KZSJfwHU8YaIRElhNRfxnKAo" http://localhost/user/1
# Respond {"message":"User 1"}
```
