package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:300;not null;" json:"first_name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updated_at"`
	Posts []Post
}
