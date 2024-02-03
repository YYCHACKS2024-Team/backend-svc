package models

type Cart struct {
	CartId            string
	No                int
	CustomerId        string
	Date              string
	MenuId            string
	Quantity          int
	Status            string
	Properties        string
	AdditionalRequest string
}
