package repository

import "go-cource-api/domain/entity"

type PostRepository interface {
	FindById(id int64) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByIdAndUserId(id int64, userId int64) (*entity.Post, error)
	Save(*entity.Post) (*entity.Post, error)
	Delete(int64) error
	Update(*entity.Post) error
}
