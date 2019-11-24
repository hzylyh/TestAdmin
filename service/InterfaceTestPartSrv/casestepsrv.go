package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/qjson"
)

var CaseStepSrv = &caseStepService{}

type caseStepService struct {
}

func (css *caseStepService) AddCaseStep(caseStepInfo *InterfaceTestPartEntity.ItfCaseStepInfo) error {
	if err := service.DB.Create(caseStepInfo).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (css *caseStepService) GetStepList(qj *qjson.QJson) (pageInfo *model.PageInfo, err error) {
	var (
		ret []*InterfaceTestPartEntity.ItfCaseStepInfo
	)

	if pageInfo, err = utils.PaginationWithDB(service.DB.Where("step_num = ?", 2), &ret, qj); err != nil {
		return nil, err
	} else {
		return pageInfo, nil
	}
}
