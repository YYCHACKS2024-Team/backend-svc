package models

import "time"

type Message struct {
	MessageId      string
	ConversationId string
	SenderId       string
	MessageText    string
	Timestamp      time.Time
}
