package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Endpoints() {
	_ = godotenv.Load()

	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.RedirectTrailingSlash = true
	
	
	router := engine.Group("/v0")
	router.POST("/webhook", func(ctx *gin.Context) {
		//recebe mensagem do whatsapp
	})
	router.GET("/webhook", func(ctx *gin.Context) {
		//Valida token da meta	
	})

	router.GET("pedidos", func(ctx *gin.Context) {
		// lista os pedidos
	})
	router.POST("pedidos", func(ctx *gin.Context) {
		// cria pedido
	})
	
	admin := router.Group("admin")
	admin.GET("estoque", func(ctx *gin.Context) {
		//retorna estoque
	})
	admin.POST("comando", func(ctx *gin.Context) {
		// envia um comando
	})
}