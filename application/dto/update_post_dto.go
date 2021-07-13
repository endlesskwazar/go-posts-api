package dto

type UpdatePostDto struct {
	Title string `validate:"required,max=255"`
	Body string `validate:"required,max=8000"`
}
