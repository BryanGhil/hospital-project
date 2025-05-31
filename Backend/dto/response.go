package dto

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ErrorRes struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type GetPageResponse struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	CountData int `json:"count_data"`
	Data      any `json:"data,omitempty"`
}
