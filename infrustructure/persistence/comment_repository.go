package persistence

import (
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"

	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepo {
	return &CommentRepo{db}
}

var _ repository.CommentRepository = &CommentRepo{}

func (r *CommentRepo) FindAll() ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.Debug().Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepo) FindById(id uint64) (*entity.Comment, error) {
	var comment entity.Comment
	err := r.db.Debug().Where("id = ?", id).Take(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepo) FindByPostId(postId uint64) ([]entity.Comment, error) {
	var comments []entity.Comment

	err := r.db.Where("post_id = ?", postId).Find(&comments).Error

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepo) Save(comment *entity.Comment) (*entity.Comment, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&comment).Error
	if err != nil {
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return comment, nil
}

func (r *CommentRepo) Delete(id uint64) error {
	err := r.db.Delete(&entity.Comment{}, id).Error

	return err
}

func (r *CommentRepo) Update(comment *entity.Comment) error {
	err := r.db.Model(&comment).Updates(comment).Error

	return err
}
