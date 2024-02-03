package roommatehandler

import "github.com/labstack/echo/v4"

type HTTPHandler struct {
	// userSvc service.User
}

func NewHTTPHandler() HTTPHandler {
	return HTTPHandler{}
}

func (hdl HTTPHandler) GetRoomById(c echo.Context) error {

	// TODO: implement here

	return nil
}

func (hdl HTTPHandler) Accept(c echo.Context) error {

	// TODO: implement here

	return nil
}

func (hdl HTTPHandler) Decine(c echo.Context) error {

	// TODO: implement here

	return nil
}
