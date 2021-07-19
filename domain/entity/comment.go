package entity

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Comment struct {
	Id        int64       `gorm:"primary_key;auto_increment" json:"id" xml:"id"`
	Body      null.String `gorm:"size:500;not null" sql:"DEFAULT:null" json:"body" xml:"body" swaggertype:"string"`
	PostId    null.Int    `gorm:"not null" json:"postId" xml:"post_id" swaggertype:"number"`
	UserId    null.Int    `gorm:"not null" json:"userId" xml:"user_id" swaggertype:"number"`
	CreatedAt time.Time   `sql:"DEFAULT:current_timestamp" json:"createAt" xml:"created_at"`
	UpdatedAt time.Time   `sql:"DEFAULT:current_timestamp" json:"updatedAt" xml:"updated_at"`
	Post      Post        `json:"-" xml:"-"`
	User      User        `json:"-" xml:"-"`
}
