package main

type Message struct {
	Data         map[string]string `json:"data,omitempty"`
	Notification *Notification     `json:"notification,omitempty"`
	Android      *AndroidConfig    `json:"android,omitempty"`
	Tokens       []string          `json:"tokens,omitempty"`
}

type Notification struct {
	Title    string `json:"title,omitempty"`
	Body     string `json:"body,omitempty"`
	ImageURL string `json:"image,omitempty"`
}

type AndroidConfig struct {
	CollapseKey string            `json:"collapse_key,omitempty"`
	Priority    string            `json:"priority,omitempty"` // one of "normal" or "high"
	Data        map[string]string `json:"data,omitempty"`     // if specified, overrides the Data field on Message type
}
