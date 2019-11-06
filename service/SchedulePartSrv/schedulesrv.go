package SchedulePartSrv

import (
	"fmt"
	"salotto/model"
	"salotto/service"
)

var Schedule = &scheService{}

type scheService struct {
}

func (srv *scheService) GetScheduleList() []*model.Schedule {
	var ret []*model.Schedule
	if err := service.DB.Find(&ret); err != nil {
		fmt.Println(err)
	}
	//service.DB.Where()
	return ret
}
