package dao

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
}
