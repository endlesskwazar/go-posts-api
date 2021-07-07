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

func (u *UserRepo) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := u.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) FindById(id uint64) (*entity.User, error) {
	var user entity.User
	err := u.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) Save(user *entity.User) (*entity.User, error) {
	err := u.db.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
