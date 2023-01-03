package telegram

import "errors"

// GetMeT struct ololo
type GetMeT struct {
	Ok     bool         `json:"ok"`
	Result GetMeResultT `json:"result"`
}

// GetUpdatesT struct ответ на getUpdates. Result = массив Result`ов
type GetUpdatesT struct {
	Ok     bool                `json:"ok"`
	Result []GetUpdatesResultT `json:"result"`
}

// GetUpdatesResultT struct ololo
type GetUpdatesResultT struct {
	UpdateID int      `json:"update_id"`
	Message  MessageT `json:"message"`
}

// MessageT struct ololo
type MessageT struct {
	MessageID int    `json:"message_id"`
	From      FromT  `json:"from"`
	Chat      ChatT  `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

// FromT struct ololo
type FromT struct {
	ID int `json:"id"`
	//IsBot        bool   `json:"is_bot"`
	FirstName   string `json:"first_name"`
	LastUpdName string `json:"lastUpd_name"`
	UserName    string `json:"username"`
	//LanguageCode string `json:"language_code"`
}

// ChatT struct ololo
type ChatT struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastUpdName string `json:"lastUpd_name"`
	UserName    string `json:"username"`
	Type        string `json:"type"`
}

// GetMeResultT struct ololo
type GetMeResultT struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	UserName  string `json:"username"`
	// CanJoinGroups bool `json:"can_join_groups"`
	// CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`
	// SupportsInlinqueries bool `json:"supports_inline_queries"`

}
type SMessageT struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

// GetSomeUpdatesT струтктура для Updates со сдвигом
type GetSomeUpdatesT struct {
	Offset  int `json:"offset"`
	Limit   int `json:"limit"`
	Timeout int `json:"timeout"`
}

type TelegramParameters struct {
	Url   string
	Token string
}

func GetParameters(url string, token string) (*TelegramParameters, error) {
	if url == "" || token == "" {
		return nil, errors.New("not enought arguments to get parameters")
	}

	p := TelegramParameters{
		Url:   url,
		Token: token,
	}
	return &p, nil
}

type TelegramCommunicator struct {
}

func NewTelegramCommunicator(params TelegramParameters) (*TelegramCommunicator, error) {
	c := TelegramCommunicator{}
	return &c, nil
}

func (t *TelegramCommunicator) ListenAndServe() {

}
