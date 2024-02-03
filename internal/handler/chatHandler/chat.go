package chathandler

import "github.com/labstack/echo/v4"

type HTTPHandler struct {
	// userSvc service.User
}

func NewHTTPHandler() HTTPHandler {
	return HTTPHandler{}
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
