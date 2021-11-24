package models

type ResponseStruct struct {
	Message     string `json:"message"`
	Error       string `json:"error"`
	Description string `json:"description"`
}
