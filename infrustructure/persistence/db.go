package persistence

import (
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repositories struct {
	Post repository.PostRepository
	User repository.UserRepository
	Comment repository.CommentRepository
	db   *gorm.DB
}

func NewRepositories(config *DatabaseConfig) (*Repositories, error) {
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Host + ")/" + config.DbName + "?parseTime=true"

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,          // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})

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
