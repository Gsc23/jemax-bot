package models

type WhatsAppWebhook struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	ID      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string     `json:"messaging_product"`
	Metadata         Metadata   `json:"metadata"`
	Contacts         []Contact  `json:"contacts"`
	Messages         []Message  `json:"messages"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID       string `json:"phone_number_id"`
}

type Contact struct {
	Profile Profile `json:"profile"`
	WaID    string  `json:"wa_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type Message struct {
	From        string              `json:"from"`
	ID          string              `json:"id"`
	Timestamp   string              `json:"timestamp"`
	Type        string              `json:"type"`
	Text        *MessageText        `json:"text,omitempty"`
	Interactive *InteractiveMessage `json:"interactive,omitempty"`
}

type InteractiveMessage struct {
	Type      string            `json:"type"`
	ListReply *ListReplyPayload `json:"list_reply,omitempty"`
}

type ListReplyPayload struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type MessageText struct {
	Body string `json:"body"`
}

type WhatsAppMessageResponse struct {
	MessagingProduct string         `json:"messaging_product"` 
	To               string         `json:"to"`                
	Type             string         `json:"type"`              
	Text             WhatsAppText   `json:"text"`
}

type WhatsAppText struct {
	Body string `json:"body"`
}
