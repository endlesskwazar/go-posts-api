package repository

import "go-cource-api/domain/entity"

type PostRepository interface {
	FindById(id int64) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByIdAndUserId(id int64, userId int64) (*entity.Post, error)
	Save(post *entity.Post) (*entity.Post, error)
	Delete(id int64) error
	Update(post *entity.Post) error
}
