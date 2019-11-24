package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
	"salotto/utils/qjson"
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

func (ms *moduleService) GetModuleList(qj *qjson.QJson) (ret []*InterfaceTestPartEntity.ProjectModule, err error) {
	if err := service.DB.Where(map[string]interface{}{"project_id": qj.GetInt("projectId"), "p_module_id": qj.GetInt("pModuleId")}).Order("created_at desc").Find(&ret).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ret, nil
}
