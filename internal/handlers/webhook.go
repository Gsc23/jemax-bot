package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsc23/jemax-bot/internal/models"
	"github.com/gsc23/jemax-bot/internal/service"
	"github.com/gsc23/jemax-bot/pkg/app"
)

func VerifyWebhook(c *gin.Context) {
	verifyToken := app.VerifyToken(c)

    mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	if mode == "subscribe" && token == verifyToken {
		c.String(http.StatusOK, challenge)
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}


func HandleWebhook(c *gin.Context) {
    var payload models.WhatsAppWebhook

    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }

    for _, entry := range payload.Entry {
        for _, change := range entry.Changes {
            if change.Value.Messages == nil {
                continue
            }

            for _, msg := range change.Value.Messages {
                phone := msg.From
                var err error

                log.Printf("Processing message from: %s", phone)

                switch msg.Type {
                case "text":
                    if msg.Text != nil {
                        err = service.ProcessSimpleText(phone, msg.Text.Body)
                    }
                case "interactive":
                    if msg.Interactive != nil && msg.Interactive.ListReply != nil {
                        id := msg.Interactive.ListReply.ID
                        err = service.ProcessMenuOption(phone, id)
                    }
                default:
                    err = service.SendErrorMenu(phone)
                }

                if err != nil {
                    log.Printf("Error processing message: %v", err)
                    _ = service.SendErrorMenu(phone)
                }
            }
        }
    }

    c.Status(http.StatusOK)
}
