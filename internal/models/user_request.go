package models

type UserGetAllResponse struct {
	CommonResponse
	Data []UserGetAllResponseBody `json:"data,omitempty"`
}

type UserGetAllResponseBody struct {
	UserId       string `json:"user_id"`
	RoleId       string `json:"role_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
}

type UserGetResponse struct {
	CommonResponse
	Data UserGetResponseBody `json:"data,omitempty"`
}

type UserGetResponseBody struct {
	UserId       string `json:"user_id"`
	RoleId       string `json:"role_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
}
