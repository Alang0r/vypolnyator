package dto

type CreateNotifyGroupDTO struct {
	UserID uint64 `json:"user_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Notifytime string `json:"notify_time"`
}
