package models

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
