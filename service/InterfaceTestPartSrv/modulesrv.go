package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/qjson"
)

var ModuleSrv = &moduleService{}

type moduleService struct {
}

func (ms *moduleService) AddProjectModule(module *InterfaceTestPartEntity.TProjectModule) error {
	module.ModuleId = utils.GenerateUUID()
	if err := service.DB.Create(module).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ms *moduleService) GetModuleList(qj *qjson.QJson) (res []map[string]interface{}, err error) {
	var ret []*InterfaceTestPartEntity.TProjectModule
	if err := service.DB.Where(map[string]interface{}{"project_id": qj.GetString("projectId"), "p_module_id": qj.GetString("pModuleId")}).Find(&ret).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	var treeData []map[string]interface{}
	for _, eachRes := range ret {
		eachMap := map[string]interface{}{
			"id":    eachRes.ModuleId,
			"value": eachRes.ModuleId,
			"label": eachRes.ModuleName,
		}
		treeData = append(treeData, eachMap)
	}
	return treeData, nil
}
