package dto

type Notification struct {
	ID            uint64 `json:"id"`
	NotifyGroupID uint64 `json:"notify_group_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Notifytime    string `json:"notify_time"`
}
