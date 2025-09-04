package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId    uint           `gorm:"primaryKey" json:"userId"`
	UserName  string         `gorm:"size:255" json:"userName"`
	Email     string         `gorm:"size:255" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "user_info"
}
