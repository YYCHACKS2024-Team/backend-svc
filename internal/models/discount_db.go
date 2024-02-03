package models

import "github.com/shopspring/decimal"

type Discount struct {
	DiscountCode string
	Type         string
	Value        decimal.Decimal
}
