package kik

import "encoding/json"

type Config struct {
	Webhook  *string  `json:"webhook"`
	Features Features `json:"features"`
}

type Features struct {
	ReceiveReadReceipts      bool `json:"receiveReadReceipts"`
	ReceiveIsTyping          bool `json:"receiveIsTyping"`
	ManuallySendReadReceipts bool `json:"manuallySendReadReceipts"`
	ReceiveDeliveryReceipts  bool `json:"receiveDeliveryReceipts"`
}

// Set bot configuration
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

// Get bot configuration
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
