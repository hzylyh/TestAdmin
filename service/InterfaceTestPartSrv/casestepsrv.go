package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/dbUtil"
	"salotto/utils/qjson"
)

var CaseStepSrv = &caseStepService{}

type caseStepService struct {
}

func (css *caseStepService) AddCaseStep(caseStepInfo *InterfaceTestPartEntity.ItfCaseStepInfo) error {
	caseStepInfo.StepId = utils.GenerateUUID()
	if err := service.DB.Create(caseStepInfo).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (css *caseStepService) GetStepList(qj *qjson.QJson) (pageInfo *model.PageInfo, err error) {
	var (
		ret []*InterfaceTestPartEntity.ItfCaseStepInfo
		//pageNum = 1.00
		//pageSize = 10.00
	)

	//if pageInfo, err = utils.PaginationWithDB(service.DB.Where("case_id = ?", qj.GetString("caseId")), &ret, qj); err != nil {
	//	return nil, err
	//} else {
	//	return pageInfo, nil
	//}

	//if err = service.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Preload("Variables").Find(&ret).Error; err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//
	//pageInfo = &model.PageInfo{
	//	PageNum:  pageNum,
	//	PageSize: pageSize,
	//	Total:    0,
	//	List:     &ret,
	//}

	hdb := dbUtil.NewHDB(service.DB, &ret)
	hdb.Paginate(qj).Preload("Variables").Find(&ret)
	pageInfo, err = hdb.Pack()
	return pageInfo, nil
}
