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
	Save(*entity.Post) (*entity.Post, map[string]string)
	FindAll() ([]entity.Post, error)
	FindById(uint64) (*entity.Post, error)
}

func (app *postApp) Save(post *entity.Post) (*entity.Post, map[string]string) {
	return app.postRepository.Save(post)
}

func (app *postApp) FindById(id uint64) (*entity.Post, error) {
	return app.postRepository.FindById(id)
}

func (app *postApp) FindAll() ([]entity.Post, error) {
	return app.postRepository.FindAll()
}
