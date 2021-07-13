package services

import (
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"
)

type postService struct {
	postRepository repository.PostRepository
}

var _ PostService = &postService{}

type PostService interface {
	Save(*entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindById(uint64) (*entity.Post, error)
	Delete(id uint64) error
	FindByIdAndUserId(id uint64, userId uint64) (*entity.Post, error)
	Update(post *entity.Post) error
}

func (s *postService) Save(post *entity.Post) (*entity.Post, error) {
	return s.postRepository.Save(post)
}

func (s *postService) FindById(id uint64) (*entity.Post, error) {
	return s.postRepository.FindById(id)
}

func (s *postService) FindAll() ([]entity.Post, error) {
	return s.postRepository.FindAll()
}

func (s *postService) FindByIdAndUserId(id uint64, userId uint64) (*entity.Post, error) {
	return s.postRepository.FindByIdAndUserId(id, userId)
}

func (s *postService) Delete(id uint64) error  {
	return s.postRepository.Delete(id)
}


func (s *postService) Update(post *entity.Post) error  {
	return s.postRepository.Update(post)
}
