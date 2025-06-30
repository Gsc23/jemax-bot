package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gsc23/jemax-bot/internal/models"
	"github.com/gsc23/jemax-bot/internal/service"
)

func VerifyWebhook(c *gin.Context) {
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	if mode == "subscribe" && token == os.Getenv("WHATSAPP_VERIFY_TOKEN") {
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
					_ = service.SendErrorMenu(phone)
				}

				if err != nil {
					_ = service.SendErrorMenu(phone)
				}
			}
		}
	}

	c.Status(http.StatusOK)
}
