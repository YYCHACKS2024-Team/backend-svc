package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	OrderId      string
	CartId       string
	CustomerId   string
	SubTotal     decimal.Decimal
	DiscountCode string
	TotalAmount  decimal.Decimal
	DateTime     time.Time
	Status       string
}
