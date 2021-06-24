package entity

import (
	"time"
)

type Post struct {
	Id        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;" json:"title"`
	Body      string    `gorm:"size:8000;not null;" json:"body"`
	UserId    uint64    `gorm:"size:100;not null;" json:"userId"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"createAt"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updatedAt"`
	User      User `json:"-"`
}
