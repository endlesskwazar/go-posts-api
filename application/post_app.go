package application

import (
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"
)

type postApp struct {
	postRepository repository.PostRepository
}

var _ PostAppInterface = &postApp{}

type PostAppInterface interface {
	Save(*entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindById(uint64) (*entity.Post, error)
	Delete(id uint64) error
	FindByIdAndUserId(id uint64, userId uint64) (*entity.Post, error)
	Update(post *entity.Post) error
}

func (app *postApp) Save(post *entity.Post) (*entity.Post, error) {
	return app.postRepository.Save(post)
}

func (app *postApp) FindById(id uint64) (*entity.Post, error) {
	return app.postRepository.FindById(id)
}

func (app *postApp) FindAll() ([]entity.Post, error) {
	return app.postRepository.FindAll()
}

func (app *postApp) FindByIdAndUserId(id uint64, userId uint64) (*entity.Post, error) {
	return app.postRepository.FindByIdAndUserId(id, userId)
}

func (app *postApp) Delete(id uint64) error  {
	return app.postRepository.Delete(id)
}


func (app *postApp) Update(post *entity.Post) error  {
	return app.postRepository.Update(post)
}
