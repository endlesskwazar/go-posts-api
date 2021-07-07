package dto

type UserDto struct {
	Name     string `validate:"required,max=255"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=40"`
}
