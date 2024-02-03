package userHandler

import (
	"net/http"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/constant"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/models"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	userSvc service.User
}

func NewHTTPHandler(userSvc service.User) HTTPHandler {
	return HTTPHandler{userSvc: userSvc}
}

func (hdl HTTPHandler) GetAll(c echo.Context) error {

	var response models.UserGetAllResponse
	resultList, err := hdl.userSvc.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	var responseBodyList = make([]models.UserGetAllResponseBody, 0, len(resultList))
	for _, each := range resultList {
		responseBodyList = append(responseBodyList, models.UserGetAllResponseBody{
			UserId:       each.UserId,
			RoleId:       each.RoleId,
			FirstName:    each.FirstName,
			LastName:     each.LastName,
			EmailAddress: each.EmailAddress,
			PhoneNumber:  each.PhoneNumber,
		})
	}

	response.Code = constant.SuccessCode
	response.Message = constant.SuccessMessage
	response.Data = responseBodyList

	return c.JSON(http.StatusOK, response)
}

func (hdl HTTPHandler) GetUserById(c echo.Context) error {

	// TODO: implement here

	return nil
}

func (hdl HTTPHandler) Register(c echo.Context) error {

	// TODO: implement here

	return nil
}

func (hdl HTTPHandler) Login(c echo.Context) error {

	// TODO: implement here

	return nil
}

func (hdl HTTPHandler) GetProfile(c echo.Context) error {

	// TODO: implement here

	return nil
}

func (hdl HTTPHandler) StorePreferences(c echo.Context) error {

	// TODO: implement here

	return nil
}
