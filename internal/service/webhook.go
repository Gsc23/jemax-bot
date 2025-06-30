package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gsc23/jemax-bot/internal/models"
	"github.com/gsc23/jemax-bot/pkg/helper"
)

var (
	WhatsAppToken = os.Getenv("WHATSAPP_ACCESS_TOKEN")
	WhatsAppPhoneID = os.Getenv("WHATSAPP_PHONE_ID")
)

func ProcessSimpleText(phone, text string) error {
	text = strings.TrimSpace(strings.ToLower(text))

	if helper.ContainsAnyWord(text, []string{"humano", "atendente", "jeanne", "max"}) {
		return sendTextMessage(phone, helper.MsgSeller)
	}


	if text != "" {
		return sendInteractiveMenu(phone)
	}

	return nil
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
			Body:   models.WhatsAppTextBlock{Text: "Como podemos te ajudar hoje?"},
			Footer: &models.WhatsAppTextBlock{Text: "Atendemos das 18h às 23h."},
			Action: models.WhatsAppListMessageAction{
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

func send(token, phoneID string, payload interface{}) error {
	url := fmt.Sprintf("https://graph.facebook.com/v18.0/%s/messages", phoneID)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		return fmt.Errorf("erro ao enviar mensagem: %s", buf.String())
	}
	return nil
}