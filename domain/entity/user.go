package entity

import (
	"time"
)

type User struct {
	Id        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:300;not null;" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updatedAt"`
	Posts []Post `json:"-"`
	Tokens []Token `json:"-"`
}
