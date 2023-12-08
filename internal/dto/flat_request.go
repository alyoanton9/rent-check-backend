package dto

type FlatRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
}
