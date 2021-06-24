package application

import (
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"
)

type userApp struct {
	userRepository repository.UserRepository
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	Save(*entity.User) (*entity.User, map[string]string)
	FindAll() ([]entity.User, error)
	FindById(uint64) (*entity.User, error)
	FindByEmail(string) (*entity.User, map[string]string)
}

func (app *userApp) Save(user *entity.User) (*entity.User, map[string]string) {
	return app.userRepository.Save(user)
}

func (app *userApp) FindById(id uint64) (*entity.User, error) {
	return app.userRepository.FindById(id)
}

func (app *userApp) FindAll() ([]entity.User, error) {
	return app.userRepository.FindAll()
}

func (app *userApp) FindByEmail(email string) (*entity.User, map[string]string) {
	return app.userRepository.FindByEmail(email)
}