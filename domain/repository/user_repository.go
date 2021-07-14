package repository

import "go-cource-api/domain/entity"

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindById(id int64) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
