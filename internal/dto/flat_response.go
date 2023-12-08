package dto

type FlatResponse struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
}
