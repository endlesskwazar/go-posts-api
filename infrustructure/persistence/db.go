package persistence

import (
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repositories struct {
	Post repository.PostRepository
	User repository.UserRepository
	Comment repository.CommentRepository
	db   *gorm.DB
}

func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	dsn := DbUser + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Post: NewPostRepository(db),
		User: NewUserRepository(db),
		Comment: NewCommentRepository(db),
		db:   db,
	}, nil
}

func (s *Repositories) Automigrate() error {
	err := s.db.AutoMigrate(&entity.Post{}, &entity.User{}, &entity.Comment{})
	if err != nil {
		return err
	}

	return nil
}
