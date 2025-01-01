package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Name        string `json:"name" gorm:"primaryKey"`
	Description string `json:"description"`
	Input       string `json:"input"`
	Output      string `json:"output"`
	Result      string `json:"result"`
}
