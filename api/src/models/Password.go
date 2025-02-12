package models

// Password represents a request format to change password
type Password struct {
	NewPassword     string `json:"novaSenha"`
	CurrentPassword string `json:"senhaAtual"`
}
