package entity

type Token struct {
	Id int64 `gorm:"primary_key;auto_increment"`
	Token string `gorm:"not null;"`
	Refresh string `gorm:"not null;"`
	UserId int64 `gorm:"not null;"`
	User User
}
