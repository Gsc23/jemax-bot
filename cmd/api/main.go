package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gsc23/jemax-bot/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env, using system environment variables")
	}

	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.RedirectTrailingSlash = true

	router := engine.Group("/v0")
	router.POST("/webhook", handlers.HandleWebhook)
	router.GET("/webhook", handlers.VerifyWebhook)

	router.GET("/pedidos", func(ctx *gin.Context) {
		// lista os pedidos
	})
	router.POST("/pedidos", func(ctx *gin.Context) {
		// cria pedido
	})

	admin := router.Group("/admin")
	admin.GET("/estoque", func(ctx *gin.Context) {
		// retorna estoque
	})
	admin.POST("/comando", func(ctx *gin.Context) {
		// envia um comando
	})

	// Sobe o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("Error runing server: %v", err)
	}
}
