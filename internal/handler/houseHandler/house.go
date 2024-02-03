package househandler

import "github.com/labstack/echo/v4"

type HTTPHandler struct {
	// userSvc service.User
}

func NewHTTPHandler() HTTPHandler {
	return HTTPHandler{}
}

func (hdl HTTPHandler) GetHouseById(c echo.Context) error {

	// TODO: implement here

	return nil
}
