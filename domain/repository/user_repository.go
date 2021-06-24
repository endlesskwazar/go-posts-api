package repository

import "go-cource-api/domain/entity"

type UserRepository interface {
	Save(*entity.User) (*entity.User, map[string]string)
	FindById(uint64) (*entity.User, error)
	FindAll() ([]entity.User, error)
	FindByEmail(email string) (*entity.User, map[string]string)
}
