package model

import "github.com/jinzhu/gorm"

type Schedule struct {
	gorm.Model
	Cron string `json:"cron"`
}
