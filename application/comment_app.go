package application

import (
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"
)

type commentApp struct {
	commentRepository repository.CommentRepository
}

var _ CommentAppInterface = &commentApp{}

type CommentAppInterface interface {
	Save(comment *entity.Comment) (*entity.Comment, map[string]string)
	FindAll() ([]entity.Comment, error)
	FindById(uint64) (*entity.Comment, error)
}

func (app *commentApp) Save(comment *entity.Comment) (*entity.Comment, map[string]string) {
	return app.commentRepository.Save(comment)
}

func (app *commentApp) FindById(id uint64) (*entity.Comment, error) {
	return app.commentRepository.FindById(id)
}

func (app *commentApp) FindAll() ([]entity.Comment, error) {
	return app.commentRepository.FindAll()
}
