package models

type User struct {
	UserId       string
	RoleId       string
	FirstName    string
	LastName     string
	EmailAddress string
	PhoneNumber  string
	PasswordHash string
}
