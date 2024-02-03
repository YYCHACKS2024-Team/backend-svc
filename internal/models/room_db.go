package models

import "github.com/shopspring/decimal"

type Room struct {
	RoomId  string
	Price   decimal.Decimal
	Address string
	Note    string
}
