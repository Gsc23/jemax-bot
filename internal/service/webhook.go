package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gsc23/jemax-bot/internal/models"
	"github.com/gsc23/jemax-bot/pkg/helper"
)

var (
	WhatsAppToken = "EAARsZAA1hFIIBPAvyytkZCG6H2D2UXUheHXGJgfSz76voCkX3ASIgkl0yshR8a7P97XS7GW1D4juCYjmAnd2JMhveZCua6NPzGZCicNPa0bf9fenrHyIZA8h0DceZCJZAFfqjY0NUb0mNwrQska7O9yn5yyhsvJV6VHcwY47acOMUO2bhMserhedPiBG5yw4rzibVAkeoof23r076OrypT2vBW4T4gWhQC7DFZBE9qYhtj7g5AZDZD"
	WhatsAppPhoneID = "671700972697735"
	apiClient = &http.Client{Timeout: 10 * time.Second} 
)

const (
    GraphAPIVersion = "v18.0"
    GraphAPIBaseURL = "https://graph.facebook.com"
)

func ProcessSimpleText(phone, text string) error {
    text = strings.TrimSpace(strings.ToLower(text))

    if helper.ContainsAnyWord(text, []string{"humano", "atendente", "falar com"}) {
        return sendTextMessage(phone, helper.MsgSeller)
    }


	if text != "" {
		return sendInteractiveMenu(phone)
	}

    return nil
}

func sendTextMessage(phone string, msg *helper.Msg) error {
	payload := models.WhatsAppMessageResponse{
		MessagingProduct: "whatsapp",
		To:               phone,
		Type:             "text",
		Text: models.WhatsAppText{
			Body: msg.Body,
		},
	}
	return send(WhatsAppToken, WhatsAppPhoneID, payload)
}

func sendInteractiveMenu(phone string) error {
	menu := models.WhatsAppListMessageRequest{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone,
		Type:             "interactive",
		Interactive: models.WhatsAppInteractiveBlock{
			Type: "list",
			Header: &models.WhatsAppHeader{Type: "text", Text: "JeMax Burger 🍔"},
			Body:   &models.WhatsAppTextBlock{Text: "Como podemos te ajudar hoje?"},
			Footer: &models.WhatsAppTextBlock{Text: "Atendemos das 18h às 23h."},
			Action: &models.WhatsAppListMessageAction{
				Button: "Ver opções",
				Sections: []models.WhatsAppListSection{
					{
						Title: "Menu Principal",
						Rows: []models.WhatsAppListRow{
							{ID: "mostrar_cardapio", Title: "Mostrar Cardápio", Description: "Exibir cardápio"},
							{ID: "fazer_pedido", Title: "Fazer Pedido", Description: "Monte seu pedido agora mesmo!"},
							{ID: "atendimento_humano", Title: "Falar com Atendente", Description: "Encaminharemos você para alguém"},
							{ID: "promocao_dia", Title: "Promoções do Dia", Description: "Exibir lista de promoções do dia!"},
						},
					},
				},
			},
		},
	}
	return send(WhatsAppToken, WhatsAppPhoneID, menu)
}

func ProcessMenuOption(phone, optionId string) error {
	switch optionId {
	case "fazer_pedido":
		return sendTextMessage(phone, helper.MsgWellcome)
	case "atendimento_humano":
		return sendTextMessage(phone, helper.MsgSeller)
	case "mostrar_cardapio":
		return sendTextMessage(phone, helper.MsgMenu)
	case "promocao_dia":
		return sendTextMessage(phone, helper.MsgPromos)
	default:
		return errors.New("opcao invalida")
	}
}

func SendErrorMenu(phone string) error {
	if err := sendTextMessage(phone, helper.MsgErrorMenu); err != nil {
		return err
	}
	return sendInteractiveMenu(phone)
}

func send(token, phoneID string, payload interface{}) error {
    url := fmt.Sprintf("%s/%s/%s/messages", GraphAPIBaseURL, GraphAPIVersion, phoneID)
    log.Printf("Sending request for: %s", url)

    jsonData, err := json.Marshal(payload)
    if err != nil {
        log.Printf("JSON Marshal Error: %v", err)
        return err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        log.Printf("Request Error: %v", err)
        return err
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)

    resp, err := apiClient.Do(req)
    if err != nil {
        log.Printf("Error at Http Call: %v", err)
        return err
    }
    defer resp.Body.Close()

    log.Printf("Meta API Response: Status %s", resp.Status)

    if resp.StatusCode >= 400 {
        buf := new(bytes.Buffer)
        buf.ReadFrom(resp.Body)
        responseBody := buf.String()
        log.Printf("Error from Meta API: %s", responseBody)
        return fmt.Errorf("error sending message: %s", responseBody)
    }

    log.Println("Message sent!")
    return nil
}