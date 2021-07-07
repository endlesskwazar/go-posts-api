package dto

type LoginUserDto struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=40"`
}