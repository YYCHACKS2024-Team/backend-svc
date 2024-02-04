package roommatehandler

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	roommateSvc service.Roommate
}

func NewHTTPHandler(roommateSvc service.Roommate) HTTPHandler {
	return HTTPHandler{roommateSvc: roommateSvc}
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
