package entity

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type User struct {
	Id        int64       `gorm:"primary_key;auto_increment" json:"id"`
	Name      null.String `gorm:"size:300;not null;" json:"name" swaggertype:"string"`
	Email     null.String `gorm:"size:100;not null;unique" json:"email" swaggertype:"string"`
	Password  null.String `gorm:"size:100;not null;" json:"password" swaggertype:"string"`
	CreatedAt time.Time   `sql:"DEFAULT:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time   `sql:"DEFAULT:current_timestamp" json:"updatedAt"`
	Posts     []Post      `json:"-"`
}
