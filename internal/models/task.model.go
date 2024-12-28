package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Id          int    `json:"id" gorm:"unique,autoincrement,primaryKey"`
	Name        string `json:"name" gorm:"primaryKey"`
	Description string `json:"description"`
	Input       string `json:"input"`
	Output      string `json:"output"`
	Result      string `json:"result"`
}
