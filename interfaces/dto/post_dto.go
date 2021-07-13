package dto

type PostDto struct {
	Title string `validate:"required,max=255" xml:"title" json:"title"`
	Body string `validate:"required,max=8000" xml:"body" json:"body"`
}