package dao

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey;autoIncrement" json:"id"`
	Name      string         `json:"name,omitempty"`
	Age       int            `json:"age,omitempty"`
	Role      string         `json:"role,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
}

type UserDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
