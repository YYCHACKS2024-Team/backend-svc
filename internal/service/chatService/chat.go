package chatservice

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
)

type chatSvc struct {
	conversationRepo database.ConversationRepository
	MessageRepo      database.MessageRepository
}

func New(conversationRepo database.ConversationRepository, messageRepo database.MessageRepository) service.Chat {
	return chatSvc{conversationRepo: conversationRepo, MessageRepo: messageRepo}
}

func (srv chatSvc) Store() error {

	//TODO: implement later

	return nil
}

func (srv chatSvc) GetCoversationById() error {

	//TODO: implement later

	return nil
}

func (srv chatSvc) GetMessageById() error {

	//TODO: implement later

	return nil
}
