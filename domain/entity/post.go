package entity

import (
	"time"
)

type Post struct {
	Id        uint64    `gorm:"primary_key;auto_increment" json:"id" xml:"id"`
	Title     string    `gorm:"size:255;not null;default:null" json:"title" xml:"title"`
	Body      string    `gorm:"size:8000;not null;default:null" json:"body" xml:"body"`
	UserId    uint64    `gorm:"not null;default:null" json:"userId" xml:"user-id"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"createAt" xml:"created-at"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updatedAt" xml:"updated-at"`
	User      User      `json:"-" xml:"-"`
	Comments  []Comment `json:"-" xml:"-"`
}
