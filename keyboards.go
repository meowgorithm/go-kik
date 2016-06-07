package kik

// Keyboard is the custom keyboard optionally included with responses
type Keyboard struct {
	To        string             `json:"to,omitempty"`
	Hidden    bool               `json:"hidden,omitempty"`
	Type      string             `json:"type,omitempty"`
	Responses []KeyboardResponse `json:"responses,omitempty"`
}

// KeyboardResponse is a button in custom keyboards
type KeyboardResponse struct {
	Type string `json:"type,omitempty"`
	Body string `json:"body,omitempty"`
}
