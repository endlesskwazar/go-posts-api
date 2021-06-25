package dto

type PostDto struct {
	Title string `validate:"required,max=255"`
	UserId uint64 `validate:"required"`
	Body string `validate:"required,max=8000"`
}