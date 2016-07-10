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

	// Pictures
	PicURL string `json:"picUrl,omitempty"`

	// Attribution
	//
	// Note that attribution values can either be a JSON object (see the
	// `Attribution` struct below) or a couple of strings:
	//
	// `gallery`, which gives it the default 'gallery' message name and icon
	// `camera` which gives it the default 'camera' message name and icon
	//
	// For more info see:
	// https://dev.kik.com/#/docs/messaging#picture
	Attribution interface{} `json:"attribution,omitempty"`
}

// Attribution storkes Kik-supported content attribution fields
type Attribution struct {
	Name    string `json:"name,omitempty"`
	IconURL string `json:"iconUrl,omitempty"`
}

// NewPictureMessage is a helper function for creating new picture message
// structs for the purpose of sending.
func NewPictureMessage(chatID *string, to *string, picURL *string, attribution interface{}) Message {
	return Message{
		Type:        Picture,
		ChatID:      *chatID,
		To:          *to,
		PicURL:      *picURL,
		Attribution: attribution,
	}
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
