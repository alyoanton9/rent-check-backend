package dto

type GroupIdFlatIdCopyRequest struct {
	GroupId    uint64  `json:"groupId"`
	FlatIdCopy *uint64 `json:"flatIdCopy"`
}
