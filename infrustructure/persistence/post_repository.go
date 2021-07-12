package persistence

import (
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"

	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepo {
	return &PostRepo{db}
}

var _ repository.PostRepository = &PostRepo{}

func (r *PostRepo) FindAll() ([]entity.Post, error) {
	var posts []entity.Post

	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepo) FindById(id uint64) (*entity.Post, error) {
	var post entity.Post
	err := r.db.Where("id = ?", id).Take(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepo) FindByIdAndUserId(id uint64, userId uint64) (*entity.Post, error) {
	var post entity.Post

	err := r.db.Where("id = ? AND user_id >= ?", id, userId).First(&post).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepo) Save(post *entity.Post) (*entity.Post, error) {
	if err := r.db.Create(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepo) Update(post *entity.Post) error {
	err := r.db.Model(&post).Updates(post).Error

	return err
}

func (r *PostRepo) Delete(id uint64) error {
	err := r.db.Delete(&entity.Post{}, id).Error

	return err
}
