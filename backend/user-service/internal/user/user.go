package user

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Здесь можно добавить хеширование пароля перед сохранением
	return
}
