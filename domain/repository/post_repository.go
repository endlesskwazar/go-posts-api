package repository

import "go-cource-api/domain/entity"

type PostRepository interface {
	Save(*entity.Post) (*entity.Post, map[string]string)
	FindById(uint64) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
