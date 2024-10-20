# Gin入門

## Ginインストール手順
```zsh
go mod init gin-fleamarket

go get -u github.com/gin-gonic/gin 
```

## main.goの作成とサンプルコード
main.go ↓
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}

```

## Goの立ち上げコマンドと動作確認コマンド
```zsh
go run main.go

curl localhost:8080/ping
```

## ホットリロードの設定と初期化と起動方法
```zsh
go install github.com/air-verse/air@latest

air init

air run main.go
```