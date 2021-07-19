package dto

type PostDto struct {
	Title string `validate:"required,max=255" xml:"title" json:"title"`
	Body  string `validate:"required,max=8000" xml:"body" json:"body"`
}

type UpdatePostDto struct {
	Title string `validate:"required,max=255" json:"title" xml:"title"`
	Body  string `validate:"required,max=8000" json:"body" xml:"body"`
}
