package entity

import (
	"time"
)

type User struct {
	Id        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:300;not null;" json:"firstName"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updated_at"`
	Posts []Post `json:"-"`
}
