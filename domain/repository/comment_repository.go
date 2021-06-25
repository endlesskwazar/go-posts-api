package repository

import "go-cource-api/domain/entity"

type CommentRepository interface {
	Save(comment *entity.Comment) (*entity.Comment, map[string]string)
	FindById(uint64) (*entity.Comment, error)
	FindAll() ([]entity.Comment, error)
}
