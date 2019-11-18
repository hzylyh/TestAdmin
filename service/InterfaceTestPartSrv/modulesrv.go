package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
)

var ModuleSrv = &moduleService{}

type moduleService struct {
}

func (ms *moduleService) AddProjectModule(module *InterfaceTestPartEntity.ProjectModule) error {
	if err := service.DB.Create(module).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
