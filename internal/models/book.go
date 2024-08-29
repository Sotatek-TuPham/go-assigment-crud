package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookName string `json:"book_name"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Shell    string `json:"shell"`
}
