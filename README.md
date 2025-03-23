# go-auth-reverse-proxy

go-auth-reverse-proxy 是一款使用 Go 語言實作的輕量級 具授權驗證機制的反向代理伺服器。
透過 Bearer Token（靜態 JSON 或 JWT）驗證機制，保護後端服務僅允許授權使用者存取。
適合用於開發、內網服務保護或作為 API Gateway 前置防線使用。

- 支援 Token 驗證（靜態或 JWT）

- 🔁 請求轉發至自訂 Proxy URL

- ⚙️ CLI 與環境變數雙重配置支援

- 🐳 支援 Docker 一鍵部署

- 📜 提供簡易 JWT 生成工具

## Quick Start

### Setting Tokens

```bash
cat data/auth-tokens.json
```

### Generate a new Token

```bash
# Create Secret Key
echo JWT_SECRET=$(openssl rand -base64 32) > .env

# Generate token
go run cmd/generate-jwt/main.go -name=<token_name>
```

## Deploy on Docker

### Windows(PowerShell)

```powershell
# http 80 port -> auth -> http 8080 port
docker run -d `
--name go-auth-reverse-proxy `
-p 80:80 `
-e PROXY_URL=http://host.docker.internal:8080 `
-v "${PWD}/data:/app/data" `
--restart unless-stopped `
alsk1369854/go-auth-reverse-proxy
```

### MacOS

```bash
# http 80 port -> auth -> http 8080 port
docker run -d \
--name go-auth-reverse-proxy \
-p 80:80 \
-e PROXY_URL=http://host.docker.internal:8080 \
-v "$(pwd)/data:/app/data" \
--restart unless-stopped \
alsk1369854/go-auth-reverse-proxy
```

### Linux

```bash
# http 80 port -> auth -> http 8080 port
docker run -d \
--name go-auth-reverse-proxy \
--network=host \
-e PROXY_URL=http://localhost:8080 \
-v "$(pwd)/data:/app/data" \
--restart unless-stopped \
alsk1369854/go-auth-reverse-proxy

```

## Track System logs

```bash
docker logs -f go-auth-reverse-proxy
```

## Dev

### Run app

```bash
go run main.go -port=80 -proxy=http://localhost:8080 -auth=./data/auth-tokens.json
```

### Test Proxy

#### Create Test Server

```bash
docker run -d \
--name py-test-server \
-p 8080:80 \
--restart unless-stopped \
alsk1369854/py-test-server
```

#### Test API

```bash
# set token
export token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NDI3NDQ4NzAsIm5hbWUiOiJhZG1pbiJ9.dnn3Cl8LwJh7fFuLufARoz1evzEBKf9Gfr3n1hHDgN0

# test api 1
curl -H "Authorization: Bearer $token" http://localhost/
# Respond {"message":"Hello world"}

# test api 2
curl -H "Authorization: Bearer $token" http://localhost/user
# Respond {"message":"User"}

# test api 3
curl -H "Authorization: Bearer $token" http://localhost/user/1
# Respond {"message":"User 1"}
```

## Docker image build and push

```bash
# image values
export tag=0.0.1
export docker_account=alsk1369854
export image_name=go-auth-reverse-proxy

# login
docker login

# build
docker build -t ${docker_account}/${image_name}:${tag} .
docker tag ${docker_account}/${image_name}:${tag} ${docker_account}/${image_name}:latest

# push to docker hub
docker push ${docker_account}/${image_name}:${tag}
docker push ${docker_account}/${image_name}:latest

# clear
docker rmi ${docker_account}/${image_name}:${tag} ${docker_account}/${image_name}:latest
```
