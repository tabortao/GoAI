package models

// GenerateRequest is the request body for the /api/v1/generate endpoint.
type GenerateRequest struct {
	Prompt   string `json:"prompt"`
	Text     string `json:"text"`
	Template string `json:"template"`
	Stream   bool   `json:"stream"`
	Model    string `json:"model"` // Optional: specify which LLM to use
}

// GenerateResponse is the non-streaming response for the /api/v1/generate endpoint.
type GenerateResponse struct {
	Text string `json:"text"`
}