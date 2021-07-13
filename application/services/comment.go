package services

import (
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"
)

type commentService struct {
	commentRepository repository.CommentRepository
}

var _ CommentService = &commentService{}

type CommentService interface {
	Save(comment *entity.Comment) (*entity.Comment, error)
	FindAll() ([]entity.Comment, error)
	FindById(uint64) (*entity.Comment, error)
	FindByPostId(postId uint64) ([]entity.Comment, error)
	Delete(uint64) error
	Update(comment *entity.Comment) error
}

func (s *commentService) Save(comment *entity.Comment) (*entity.Comment, error) {
	return s.commentRepository.Save(comment)
}

func (s *commentService) FindById(id uint64) (*entity.Comment, error) {
	return s.commentRepository.FindById(id)
}

func (s *commentService) FindAll() ([]entity.Comment, error) {
	return s.commentRepository.FindAll()
}

func (s *commentService) FindByPostId(postId uint64) ([]entity.Comment, error) {
	return s.commentRepository.FindByPostId(postId)
}

func (s *commentService) Delete(id uint64) error {
	return s.commentRepository.Delete(id)
}

func (s *commentService) Update(comment *entity.Comment) error {
	return s.commentRepository.Update(comment)
}
