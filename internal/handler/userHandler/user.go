package carthandler

import (
	"net/http"

	apiresponse "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler/api_response"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/models"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
}

func NewHTTPHandler() HTTPHandler {
	return HTTPHandler{}
}

func (hdl HTTPHandler) GetAllUsers(c echo.Context) error {

	var req models.AddToCartRequest
	var res models.AddToCartResponse

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, apiresponse.InvalidInputError(nil))
	}

	// cartId, no, err := hdl.cartService.AddToCart(req)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, apiresponse.InternalError(err))
	// }

	// res.CommonResponse = apiresponse.SuccessResponse()
	// res.Data.CartId = cartId
	// res.Data.No = no

	return c.JSON(http.StatusOK, res)
}
