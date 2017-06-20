package dto

// UserDTO user dto represent model
type UserDTO struct {
	Name string `json:"user_name"`
}

// UserRegisterDTO dto for users/register
type UserRegisterDTO struct {
	LatName     string `json:"last_name"`
	FirstName   string `json:"first_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
