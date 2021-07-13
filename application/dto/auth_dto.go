package dto

type LoginUserDto struct {
	Email    string `validate:"required,email" json:"email" xml:"email"`
	Password string `validate:"required,min=8,max=40" json:"password" xml:"password"`
}

type RegisterUserDto struct {
	Name     string `validate:"required,max=255" json:"name" xml:"name"`
	Email    string `validate:"required,email" json:"email" xml:"email"`
	Password string `validate:"required,min=8,max=40" json:"password" xml:"password"`
}
