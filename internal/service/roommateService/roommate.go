package roommateservice

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
)

type roommateSvc struct {
	userRepo      database.UserRepository
	roomRepo      database.RooomRepository
	matchHistRepo database.MatchHistRepository
}

func New(userRepo database.UserRepository, roomRepo database.RooomRepository, matchHistRepo database.MatchHistRepository) service.Roommate {
	return roommateSvc{userRepo: userRepo, roomRepo: roomRepo, matchHistRepo: matchHistRepo}
}

func (srv roommateSvc) Accept() error {

	//TODO: implement later

	return nil
}

func (srv roommateSvc) Decline() error {

	//TODO: implement later

	return nil
}
