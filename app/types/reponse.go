package types

type APIResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Data       string `json:"data,omitempty"`
}
