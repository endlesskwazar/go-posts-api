package repository

import "go-cource-api/domain/entity"

type PostRepository interface {
	FindById(uint64) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByIdAndUserId(uint64, uint64) (*entity.Post, error)
	Save(*entity.Post) (*entity.Post, map[string]string)
	Delete(uint64) error
	Update(*entity.Post) error
}
