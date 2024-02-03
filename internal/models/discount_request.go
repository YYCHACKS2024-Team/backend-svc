package models

import (
	"github.com/shopspring/decimal"
)

type GetDiscountByCodeResponse struct {
	CommonResponse
	Data GetDiscountByCodeResponseBody `json:"data,omitempty"`
}

type GetDiscountByCodeResponseBody struct {
	Type  string          `json:"type"`
	Value decimal.Decimal `json:"value"`
}
