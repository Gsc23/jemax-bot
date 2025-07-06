package service

import (
	"errors"
	"log"

	"github.com/gsc23/jemax-bot/internal/models"
	"github.com/gsc23/jemax-bot/pkg/helper"
)

func ProcessOrderStep(phone, text string, state string) error {
	return nil
}

func startOrder(phone string) error {
	return sendBurgerMenu(phone)
}

func sendBurgerMenu(phone string) error {
	burgerMenu := models.WhatsAppListMessageRequest{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone,
		Type:             "interactive",
		Interactive: models.WhatsAppInteractiveBlock{
			Type:   "list",
			Header: &models.WhatsAppHeader{Type: "text", Text: "Selecione seu Lanche"},
			Body:   &models.WhatsAppTextBlock{Text: "Selecione seu lanche na lista abaixo:"},
			Footer: &models.WhatsAppTextBlock{Text: "Todos os lanches!!"},
			Action: &models.WhatsAppListMessageAction{
				Button: "Ver opções",
				Sections: []models.WhatsAppListSection{
					{
						Title: "Hambúrgueres:",
						Rows: []models.WhatsAppListRow{
							{ID: "jemax_burger", Title: "Jemax Burger", Description: "Pão de Brioche, Carne Angus, Bacon e Cheddar."},
							{ID: "jemax_2.0", Title: "Jemax Burger 2.0", Description: "Pão de Brioche, Duas Carnes, Bacon e Cheddar."},
							{ID: "max_burger", Title: "Max Burger", Description: "Pão de Gergelim, Carne Angus, Mussarela e Molho Especial."},
							{ID: "duplo_max", Title: "Duplo Max", Description: "Pão de Gergelim, Duas Carnes, Mussarela e Molho Especial."},
							{ID: "je_australiano", Title: "Je Australiano", Description: "Pão Australiano, Carne Angus, Linguiça Calabresa e Cheddar."},
							{ID: "je_australiano_2.0", Title: "Je Australiano 2.0", Description: "Pão Australiano, Duas Carnes, Linguiça Calabresa e Cheddar."},
							{ID: "max_onion", Title: "Max Onion", Description: "Pão de Brioche, Carne Angus, Cheddar e Onion Rings."},
							{ID: "max_onion_2.0", Title: "Max Onion 2.0", Description: "Pão de Brioche, Duas Carnes, Cheddar e Onion Rings."},
							{ID: "max_chicken", Title: "Max Chicken", Description: "Pão de Gergelim, Frango Empanado, Cheddar e Bacon."},
						},
					},
				},
			},
		},
	}
	return send(WhatsAppToken, WhatsAppPhoneID, burgerMenu)
}

func ProcessOrderMenuOption(phone, optionId string) error {
	switch optionId {
	case "fazer_pedido":
		if err := sendTextMessage(phone, helper.MsgWellcome); err != nil {
			log.Printf("Error sending wellcome message for %s: %v", phone, err)
			return err
		}
		// return startOrder(phone)
	case "atendimento_humano":
		return sendTextMessage(phone, helper.MsgSeller)
	case "mostrar_cardapio":
		return sendTextMessage(phone, helper.MsgMenu)
	case "promocao_dia":
		return sendTextMessage(phone, helper.MsgPromos)
	default:
		return errors.New("opcao invalida")
	}

	return  nil
}