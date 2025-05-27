package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORSの設定
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// ヘルスチェックエンドポイント
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// ブログ記事取得エンドポイント
	r.GET("/api/blogs", func(c *gin.Context) {
		// TODO: AWSからブログ記事を取得する処理を実装
		c.JSON(http.StatusOK, gin.H{
			"message": "ブログ記事一覧を取得するエンドポイント",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("サーバーの起動に失敗しました:", err)
	}
} 