package model

import "time"

// CreateChatRequest is service layer create chat representation.
type CreateChatRequest struct {
	UserIDs []int
}

// CreateChatResponse is service layer chat representation.
type CreateChatResponse struct {
	UserIDs []int
	ChatID  int
}

// SendMessageRequest is service layer message representation.
type SendMessageRequest struct {
	ChatID      int
	FromUserID  int
	MessageText string
	Timestamp   time.Time
}
