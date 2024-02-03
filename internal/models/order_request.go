package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateNewOrderRequest struct {
	CartId       string `json:"cart_id,omitempty"`
	DiscountCode string `json:"discount_code,omitempty"`
}

type CreateNewOrderResponse struct {
	CommonResponse
	Data CreateNewOrderRequestBody `json:"data,omitempty"`
}

type CreateNewOrderRequestBody struct {
	OrderId string
}

type GetOrderByCustomerIdResponse struct {
	CommonResponse
	Data []GetOrderByCustomerIdResponseBody `json:"data,omitempty"`
}

type GetOrderByCustomerIdResponseBody struct {
	OrderId      string                           `json:"order_id"`
	CartId       string                           `json:"cart_id"`
	CartDetail   []GetOrderByCustomerIdCartDetail `json:"cart_detail"`
	CustomerId   string                           `json:"customer_id"`
	SubTotal     decimal.Decimal                  `json:"sub_total"`
	DiscountCode string                           `json:"discount_code,omitempty"`
	TotalAmount  decimal.Decimal                  `json:"total_amount"`
	DateTime     time.Time                        `json:"date_time"`
	Status       string                           `json:"status"`
}

type GetOrderByCustomerIdCartDetail struct {
	No                int             `json:"no"`
	CustomerId        string          `json:"customer_id"`
	Date              string          `json:"date"`
	MenuId            string          `json:"menu_id"`
	MenuName          string          `json:"menu_name"`
	Quantity          int             `json:"quantity"`
	Status            string          `json:"status"`
	Properties        string          `json:"properties,omitempty"`
	AdditionalRequest string          `json:"additional_request,omitempty"`
	Price             decimal.Decimal `json:"price"`
}

type GetOrderByDateResponse struct {
	CommonResponse
	Data []GetOrderByDateResponseBody `json:"data,omitempty"`
}

type GetOrderByDateResponseBody struct {
	OrderId      string          `json:"order_id"`
	CartId       string          `json:"cart_id"`
	CustomerId   string          `json:"customer_id"`
	SubTotal     decimal.Decimal `json:"sub_total"`
	DiscountCode string          `json:"discount_code,omitempty"`
	TotalAmount  decimal.Decimal `json:"total_amount"`
	DateTime     time.Time       `json:"date_time"`
	Status       string          `json:"status"`
}

type GetActiveOrdersResponse struct {
	CommonResponse
	Data []GetActiveOrdersResponseBody `json:"data,omitempty"`
}

type GetActiveOrdersResponseBody struct {
	OrderId      string          `json:"order_id,omitempty"`
	CustomerId   string          `json:"customer_id,omitempty"`
	CartId       string          `json:"cart_id,omitempty"`
	CartDetail   []CartDetail    `json:"cart_detail,omitempty"`
	SubTotal     decimal.Decimal `json:"sub_total,omitempty"`
	DiscountCode string          `json:"discount_code,omitempty"`
	TotalAmount  decimal.Decimal `json:"total_amount,omitempty"`
	DateTime     time.Time       `json:"date_time,omitempty"`
	Status       string          `json:"status,omitempty"`
}

type CartDetail struct {
	CartId            string               `json:"cart_id,omitempty"`
	No                int                  `json:"no,omitempty"`
	Date              string               `json:"date,omitempty"`
	MenuId            string               `json:"menu_id,omitempty"`
	MenuName          string               `json:"menu_name,omitempty"`
	Price             float64              `json:"price,omitempty"`
	Quantity          int                  `json:"quantity,omitempty"`
	Status            string               `json:"status,omitempty"`
	Properties        CartDetailProperties `json:"properties,omitempty"`
	AdditionalRequest string               `json:"additional_request,omitempty"`
}

type CartDetailProperties struct {
	Type    string `json:"type,omitempty"`
	Size    string `json:"size,omitempty"`
	Topping string `json:"toppings,omitempty"`
}
