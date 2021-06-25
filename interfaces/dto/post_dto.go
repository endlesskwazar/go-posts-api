package dto

type PostDto struct {
	Title string `validate:"required"`
	UserId uint64 `validate:"required"`
}