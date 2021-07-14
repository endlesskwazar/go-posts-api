package entity

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Post struct {
	Id        int64    `gorm:"primary_key;auto_increment" json:"id" xml:"id"`
	Title     null.String    `gorm:"size:255;not null" json:"title" xml:"title" swaggertype:"string"`
	Body      null.String    `gorm:"size:8000;not null" json:"body" xml:"body" swaggertype:"string"`
	UserId    null.Int    `gorm:"not null" json:"userId" xml:"user_id" swaggertype:"number"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"createAt" xml:"created_at"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updatedAt" xml:"updated_at"`
	User      User      `json:"-" xml:"-"`
	Comments  []Comment `json:"-" xml:"-"`
}
