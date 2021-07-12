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
	Save(comment *entity.Comment) (*entity.Comment, error)
	FindAll() ([]entity.Comment, error)
	FindById(uint64) (*entity.Comment, error)
	FindByPostId(postId uint64) ([]entity.Comment, error)
	Delete(uint64) error
	Update(comment *entity.Comment) error
}

func (app *commentApp) Save(comment *entity.Comment) (*entity.Comment, error) {
	return app.commentRepository.Save(comment)
}

func (app *commentApp) FindById(id uint64) (*entity.Comment, error) {
	return app.commentRepository.FindById(id)
}

func (app *commentApp) FindAll() ([]entity.Comment, error) {
	return app.commentRepository.FindAll()
}

func (app *commentApp) FindByPostId(postId uint64) ([]entity.Comment, error) {
	return app.commentRepository.FindByPostId(postId)
}

func (app *commentApp) Delete(id uint64) error {
	return app.commentRepository.Delete(id)
}

func (app *commentApp) Update(comment *entity.Comment) error {
	return app.commentRepository.Update(comment)
}
