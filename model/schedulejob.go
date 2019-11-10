package model

import "github.com/jinzhu/gorm"

type ScheduleJob struct {
	gorm.Model
	JobId   int
	TaskId  int
	Command string
	Depends string
}
