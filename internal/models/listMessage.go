package models

	type WhatsAppListMessageRequest struct {
		MessagingProduct string                   `json:"messaging_product"`
		RecipientType    string                   `json:"recipient_type"`   
		To               string                   `json:"to"`               
		Type             string                   `json:"type"`             
		Interactive      WhatsAppInteractiveBlock `json:"interactive"`
	}

	type WhatsAppInteractiveBlock struct {
		Type   string                     `json:"type"`
		Header *WhatsAppHeader            `json:"header,omitempty"`
		Body   *WhatsAppTextBlock          `json:"body"`
		Footer *WhatsAppTextBlock         `json:"footer,omitempty"`
		Action *WhatsAppListMessageAction  `json:"action"`
	}

	type WhatsAppHeader struct {
		Type string `json:"type"` 
		Text string `json:"text"`
	}

	type WhatsAppTextBlock struct {
		Text string `json:"text"`
	}

	type WhatsAppListMessageAction struct {
		Button   string                `json:"button"`
		Sections []WhatsAppListSection `json:"sections"`
	}

	type WhatsAppListSection struct {
		Title string            `json:"title"`
		Rows  []WhatsAppListRow `json:"rows"`
	}

	type WhatsAppListRow struct {
		ID          string `json:"id"`          
		Title       string `json:"title"`       
		Description string `json:"description"` 
	}
