package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gsc23/jemax-bot/internal/handlers"
)

func Initrouter(r *gin.RouterGroup) {
	r.POST("/webhook", handlers.HandleWebhook)
	r.GET("/webhook", handlers.VerifyWebhook)

	r.GET("/pedidos", func(ctx *gin.Context) {
		// lista os pedidos
	})
	r.POST("/pedidos", func(ctx *gin.Context) {
		// cria pedido
	})

	admin := r.Group("/admin")
	admin.GET("/estoque", func(ctx *gin.Context) {
		// retorna estoque
	})
	admin.POST("/comando", func(ctx *gin.Context) {
		// envia um comando
	})
}
