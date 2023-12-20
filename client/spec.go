package client

type Spec struct {
	Token     string `json:"token"`
	BaseURL   string `json:"base_url,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
}
