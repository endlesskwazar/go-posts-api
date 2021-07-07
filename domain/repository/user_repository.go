package repository

import "go-cource-api/domain/entity"

type UserRepository interface {
	Save(*entity.User) (*entity.User, error)
	FindById(uint64) (*entity.User, error)
	FindAll() ([]entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
