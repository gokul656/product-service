package types

type APIResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Data       string `json:"data,omitempty"`
}
