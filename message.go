package kik

import "encoding/json"

// Type for both inbound and outbound messages
type Message struct {
	ChatId               string     `json:"chatId,omitempty"`
	Id                   string     `json:"id,omitempty"`
	Type                 string     `json:"type,omitempty"`
	To                   string     `json:"to,omitempty"`
	From                 string     `json:"from,omitempty"`
	Participants         []string   `json:"participants,omitempty"`
	Body                 string     `json:"body,omitempty"`
	Timestamp            int64      `json:"timestamp,omitempty"`
	ReadReceiptRequested bool       `json:"readReceiptRequested,omitempty"`
	Mention              *string    `json:"mentions,omitempty"`
	Keyboards            []Keyboard `json:"keyboards,omitempty"`
}

// Custom keyboard type
type Keyboard struct {
	To        string             `json:"to,omitempty"`
	Hidden    bool               `json:"hidden,omitempty"`
	Type      string             `json:"type,omitempty"`
	Responses []KeyboardResponse `json:"responses,omitempty"`
}

// Button value type in custom keyboards
type KeyboardResponse struct {
	Type string `json:"type,omitempty"`
	Body string `json:"body,omitempty"`
}

// Send messages to Kik
//
// POST /message
func (c *Client) SendMessages(m []Message) error {
	payload := struct {
		Messages []Message `json:"messages"`
	}{m}
	serialized, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	data := []byte(serialized)
	_, _, err = c.apiRequest(post, "/message", &data)
	return err
}
