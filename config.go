package kik

import "encoding/json"

// Config is used to configure the Kik client on the server
type Config struct {
	Webhook  *string  `json:"webhook"`
	Features Features `json:"features"`
}

// Features contains Kik bot configuration features
type Features struct {
	ReceiveReadReceipts      bool `json:"receiveReadReceipts"`
	ReceiveIsTyping          bool `json:"receiveIsTyping"`
	ManuallySendReadReceipts bool `json:"manuallySendReadReceipts"`
	ReceiveDeliveryReceipts  bool `json:"receiveDeliveryReceipts"`
}

// SetConfig sets the bot configuration
//
// POST /config
func (c *Client) SetConfig(config Config) error {
	payload, err := json.Marshal(config)
	if err != nil {
		return err
	}
	data := []byte(payload)
	_, _, err = c.apiRequest(post, "/config", &data)
	return err
}

// GetConfig returns the bot configuration
//
// GET /config
func (c *Client) GetConfig() (conf Config, err error) {
	var payload []byte
	if _, payload, err = c.apiRequest(get, "/config", nil); err != nil {
		return conf, err
	}
	err = json.Unmarshal(payload, &conf)
	return conf, err
}
