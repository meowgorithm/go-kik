package kik

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
