package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gsc23/jemax-bot/internal/router"
	"github.com/gsc23/jemax-bot/pkg/app"
)

func main() {
	ctx := context.Background()
	mainApp := app.NewApp(ctx)
	
	if mainApp == nil {
        fmt.Println("Falha ao inicializar a aplicação. Verifique os logs de erro.")
        return
    }

	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.RedirectTrailingSlash = true
	group := engine.Group("/v0")
	
	router.Initrouter(group)

	mainApp.Start(ctx)
	_ = engine.Run(fmt.Sprintf(":%d", mainApp.Config.ServerPort))
}
