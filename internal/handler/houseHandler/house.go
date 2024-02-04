package househandler

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	houseSvc service.House
}

func NewHTTPHandler(houseSvc service.House) HTTPHandler {
	return HTTPHandler{houseSvc: houseSvc}
}

func (hdl HTTPHandler) GetHouseById(c echo.Context) error {

	// TODO: implement here

	return nil
}
