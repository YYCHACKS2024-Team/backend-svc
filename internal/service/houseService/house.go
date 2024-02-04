package houseservice

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
)

type houseSvc struct {
	roomRepo database.RooomRepository
}

func New(roomRepo database.RooomRepository) service.House {
	return houseSvc{roomRepo: roomRepo}
}

func (srv houseSvc) GetHouseById() error {

	//TODO: implement later

	return nil
}
