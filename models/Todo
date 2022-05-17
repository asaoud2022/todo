package models

import "gorm.io/gorm"

// Todo is a struct holding the todos settings.
type Todo struct {
	gorm.Model
	Id        int    `gorm:"primaryKey"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
