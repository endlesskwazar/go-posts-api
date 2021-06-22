package entity

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Id        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;" json:"title"`
	Body      string    `gorm:"size:8000;not null;" json:"body"`
	UserId    uint64    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"created_at"`
	User      User
}
