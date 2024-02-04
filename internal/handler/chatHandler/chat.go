package chathandler

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	chatSvc service.Chat
}

func NewHTTPHandler(chatSvc service.Chat) HTTPHandler {
	return HTTPHandler{chatSvc: chatSvc}
}

func (hdl HTTPHandler) StoreChatHistory(c echo.Context) error {

	// TODO: implement here

	return nil
}

func (hdl HTTPHandler) GetChatByConversationId(c echo.Context) error {

	// TODO: implement here

	return nil
}

func (hdl HTTPHandler) GetChatByMessageId(c echo.Context) error {

	// TODO: implement here

	return nil
}
