# --- 階段 1: 建構階段 ---
FROM golang:1.24.1 AS builder

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 並先下載依賴（增加快取機會）
COPY go.mod go.sum ./
RUN go mod download

# 複製整個專案
COPY . .

# 編譯 main.go，並輸出為 app
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./main.go

# --- 階段 2: 執行階段 ---
FROM debian:bookworm-slim

# 複製執行檔與 JSON 檔
WORKDIR /app
COPY --from=builder /app/app .
COPY auth-tokens.json .

# 設定預設 port
EXPOSE 80

# 預設執行 command，可被 docker run 時覆寫
ENV PROXY_URL=http://host.docker.internal:8080

CMD ["sh", "-c", "./app -port=80 -proxy=${PROXY_URL} -auth=auth-tokens.json"]
