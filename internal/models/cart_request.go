package models

import "github.com/shopspring/decimal"

type AddToCartRequest struct {
	CustomerId        string `json:"customer_id,omitempty"`
	MenuId            string `json:"menu_id,omitempty"`
	Quantity          int    `json:"quantity,omitempty"`
	Properties        string `json:"properties,omitempty"`
	AdditionalRequest string `json:"additional_request,omitempty"`
}

type AddToCartResponse struct {
	CommonResponse
	Data AddToCartResponseBody `json:"data,omitempty"`
}

type AddToCartResponseBody struct {
	CartId string `json:"cart_id,omitempty"`
	No     int    `json:"no,omitempty"`
}

type GetCartByCustomerIdResponse struct {
	CommonResponse
	Data GetCartByCustomerIdResponseBody `json:"data,omitempty"`
}

type GetCartByCustomerIdResponseBody struct {
	CartId     string                          `json:"cart_id,omitempty"`
	CartDetail []GetCartByCustomerIdCartDetail `json:"cart_detail,omitempty"`
}
type GetCartByCustomerIdCartDetail struct {
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

type GetCartByIdResponse struct {
	CommonResponse
	Data []GetCartByIdResponseBody `json:"data,omitempty"`
}

type GetCartByIdResponseBody struct {
	CartId            string `json:"cart_id"`
	No                int    `json:"no"`
	CustomerId        string `json:"customer_id"`
	Date              string `json:"date"`
	MenuId            string `json:"menu_id"`
	Quantity          int    `json:"quantity"`
	Status            string `json:"status"`
	Properties        string `json:"properties,omitempty"`
	AdditionalRequest string `json:"additional_request,omitempty"`
}
