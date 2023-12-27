package dto

type GroupItemsResponse struct {
	GroupId uint64           `json:"groupId"`
	Items   []ItemWithStatus `json:"items"`
}

type ItemWithStatus struct {
	Item   ItemResponse `json:"item"`
	Status string       `json:"status"`
}
