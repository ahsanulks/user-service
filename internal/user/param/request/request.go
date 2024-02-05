package request

type CreateUser struct {
	PhoneNumber string
	FullName    string
	Password    string
}

type UpdateProfile struct {
	PhoneNumber *string
	FullName    *string
}

type GenerateUserTokenRequest struct {
	PhoneNumber string
	Password    string
}
