package repository

import "go-cource-api/domain/entity"

type CommentRepository interface {
	Save(comment *entity.Comment) (*entity.Comment, error)
	FindByPostId(id int64) ([]entity.Comment, error)
	FindById(id int64) (*entity.Comment, error)
	FindAll() ([]entity.Comment, error)
	Delete(id int64) error
	Update(comment *entity.Comment) error
}

