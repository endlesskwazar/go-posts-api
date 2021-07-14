package entity

import (
	"time"
)

type Post struct {
	Id        uint64    `gorm:"primary_key;auto_increment" json:"id" xml:"id"`
	Title     string    `gorm:"size:255;not null" sql:"DEFAULT:null" json:"title" xml:"title"`
	Body      string    `gorm:"size:8000;not null" sql:"DEFAULT:null" json:"body" xml:"body"`
	UserId    uint64    `gorm:"not null" sql:"DEFAULT:null" json:"userId" xml:"user_id"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"createAt" xml:"created_at"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updatedAt" xml:"updated_at"`
	User      User      `json:"-" xml:"-"`
	Comments  []Comment `json:"-" xml:"-"`
}
