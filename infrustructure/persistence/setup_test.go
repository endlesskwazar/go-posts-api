package persistence

import (
	"go-cource-api/domain/entity"
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
		Name: "test",
		Password: "$2y$14$1ydnM3J094ycL7Fe6CSxT.y6O8airXr.sdlUq0.MMYHCONIMzkdv6",
		Email: "test@mail.com",
	}

	db.Create(&user)

	return user
}

func SeedPost(db *gorm.DB) *entity.Post {
	user := SeedUser(db)

	post := &entity.Post{
		Title: "Test post",
		Body: "Test body",
		UserId: user.Id,
	}

	db.Create(&post)

	return post
}

func setUpTestDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = db.Debug().AutoMigrate(
		entity.User{},
		entity.Comment{},
		entity.Post{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
