package dto

type CommentDto struct {
	Body string `validation:"required,max:500" json:"body" xml:"body"`
}

type UpdateCommentDto struct {
	Body string `validation:"required,max:500" json:"body" xml:"body"`
}
