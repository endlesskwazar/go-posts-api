package persistence

import (
	"go-cource-api/domain/entity"
	"gopkg.in/guregu/null.v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBConn() *gorm.DB {
	conn, err := setUpTestDb()

	if err != nil {
		panic(err)
	}

	return conn
}

func SeedUser(db *gorm.DB) *entity.User {
	user := &entity.User{
		Name:     null.StringFrom("test"),
		Password: null.StringFrom("$2y$14$1ydnM3J094ycL7Fe6CSxT.y6O8airXr.sdlUq0.MMYHCONIMzkdv6"),
		Email:    null.StringFrom("test@mail.com"),
	}

	db.Create(&user)

	return user
}

func SeedPost(db *gorm.DB) *entity.Post {
	user := SeedUser(db)

	post := &entity.Post{
		Title:  null.StringFrom("Test post"),
		Body:   null.StringFrom("Test body"),
		UserId: null.IntFrom(user.Id),
	}

	db.Create(&post)

	return post
}

func SeedComment(db *gorm.DB) *entity.Comment {
	post := SeedPost(db)

	comment := &entity.Comment{
		Body:   null.StringFrom("Test body"),
		UserId: post.UserId,
		PostId: null.IntFrom(post.Id),
	}

	db.Create(comment)

	return comment
}

func setUpTestDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		entity.User{},
		entity.Comment{},
		entity.Post{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
