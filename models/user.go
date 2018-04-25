package models

import "github.com/jinzhu/gorm"

// User clase usuario
type User struct {
	gorm.Model
	Username string
}
