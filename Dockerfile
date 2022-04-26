# ビルドイメージの指定
FROM golang:1.16 as builder
# 環境変数の設定
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src
COPY ./ ./
RUN go mod tidy
RUN go build -o main

# マルチステージビルド
FROM alpine
# builderからapp
COPY --from=builder /go/src /app
EXPOSE 8080
CMD ["./app/main"]
