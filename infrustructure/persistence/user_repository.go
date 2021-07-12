package persistence

import (
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

var _ repository.UserRepository = &UserRepo{}

func (u *UserRepo) FindById(id uint64) (*entity.User, error) {
	var user entity.User

	if err := u.db.Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) Save(user *entity.User) (*entity.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}

	if err := u.db.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
