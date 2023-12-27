package dto

type ItemStatusRequest struct {
	FlatId  uint64 `json:"flatId"`
	GroupId uint64 `json:"groupId"`
	ItemId  uint64 `json:"itemId"`
	Status  string `json:"status"`
}
