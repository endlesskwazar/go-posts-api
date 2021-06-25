package entity

import (
	"time"
)

type Comment struct {
	Id        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Body      string    `gorm:"size:500;not null;" json:"body"`
	PostId    uint64    `gorm:"not null;" json:"postId"`
	UserId    uint64    `gorm:"not null" json:"userId"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"createAt"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updatedAt"`
	Post      Post      `json:"-"`
	User      User      `json:"-"`
}
