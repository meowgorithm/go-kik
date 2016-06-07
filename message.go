package kik

import "encoding/json"

// Message is used to serialize and unserialize both inbound and outbound
// messages alike
type Message struct {
	ChatID               string     `json:"chatId,omitempty"`
	ID                   string     `json:"id,omitempty"`
	Type                 string     `json:"type,omitempty"`
	To                   string     `json:"to,omitempty"`
	From                 string     `json:"from,omitempty"`
	Participants         []string   `json:"participants,omitempty"`
	Body                 string     `json:"body,omitempty"`
	Timestamp            int64      `json:"timestamp,omitempty"`
	ReadReceiptRequested bool       `json:"readReceiptRequested,omitempty"`
	Mention              *string    `json:"mentions,omitempty"`
	TypeTime             int        `json:"typeTime,omitempty"`
	Keyboards            []Keyboard `json:"keyboards,omitempty"`
}

// SendMessages sends a slice of messages to Kik
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
