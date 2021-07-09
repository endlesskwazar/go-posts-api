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
	err := r.db.Debug().Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepo) FindById(id uint64) (*entity.Post, error) {
	var post entity.Post
	err := r.db.Debug().Where("id = ?", id).Take(&post).Error
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

func (r *PostRepo) Save(post *entity.Post) (*entity.Post, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&post).Error
	if err != nil {
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
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
