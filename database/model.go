package database

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name   string `gorm:"not null;" json:"name"`
	Number string `gorm:"unique; not null;" json:"number"`
}
