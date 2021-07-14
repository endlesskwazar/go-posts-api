package entity

import (
	"time"
)

type Comment struct {
	Id        uint64    `gorm:"primary_key;auto_increment" json:"id" xml:"id"`
	Body      string    `gorm:"size:500;not null" sql:"DEFAULT:null" json:"body" xml:"body"`
	PostId    uint64    `gorm:"not null" sql:"DEFAULT:null" json:"postId" xml:"post_id"`
	UserId    uint64    `gorm:"not null" sql:"DEFAULT:null" json:"userId" xml:"user_id"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"createAt" xml:"created_at"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updatedAt" xml:"updated_at"`
	Post      Post      `json:"-" xml:"-"`
	User      User      `json:"-" xml:"-"`
}
