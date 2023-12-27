package entity

type GroupItem struct {
	Id          uint64
	UserId      string
	Title       string
	Description string
	Hide        bool
	Status      string
	GroupId     uint64
}
