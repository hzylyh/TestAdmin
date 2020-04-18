package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/model/vo"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/dbUtil"
	"salotto/utils/qjson"
)

var CaseStepSrv = &caseStepService{}

type caseStepService struct {
}

func (css *caseStepService) AddCaseStep(caseStepInfo *vo.CaseStepInfoVO) error {
	stepId := utils.GenerateUUID()
	tx := service.DB.Begin()
	newStepInfo := &InterfaceTestPartEntity.TItfCaseStepInfo{
		CaseId:   caseStepInfo.CaseId,
		ItfId:    caseStepInfo.ItfId,
		StepId:   stepId,
		StepNum:  caseStepInfo.StepNum,
		StepName: caseStepInfo.StepName,
		StepDesc: caseStepInfo.StepDesc,
		ReqData:  caseStepInfo.ReqData,
		ExpRes:   caseStepInfo.ExpRes,
	}
	//newAssertInfo := &InterfaceTestPartEntity.TAssertInfo{
	//	StepId:    stepId,
	//	AssertCol: caseStepInfo.AssertCol,
	//	Method:    caseStepInfo.Method,
	//	ExpValue:  caseStepInfo.ExpValue,
	//}
	newAssertInfos := caseStepInfo.AssertInfos
	newVarInfos := caseStepInfo.Variables
	//newVarInfo := &InterfaceTestPartEntity.TCaseStepVarInfo{
	//	StepId:          stepId,
	//	CollectCol:      caseStepInfo.CollectCol,
	//	CollectColAlias: caseStepInfo.CollectColAlias,
	//}
	if err := tx.Create(newStepInfo).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	for _, assertInfo := range newAssertInfos {
		newAssertInfo := &InterfaceTestPartEntity.TAssertInfo{
			StepId:    stepId,
			AssertCol: assertInfo.AssertCol,
			Method:    assertInfo.Method,
			ExpValue:  assertInfo.ExpValue,
		}
		if err := tx.Create(newAssertInfo).Error; err != nil {
			fmt.Println(err)
			tx.Rollback()
			return err
		}
	}

	for _, varInfo := range newVarInfos {
		newVarInfo := &InterfaceTestPartEntity.TCaseStepVarInfo{
			StepId:          stepId,
			CollectCol:      varInfo.CollectCol,
			CollectColAlias: varInfo.CollectColAlias,
		}
		if err := tx.Create(newVarInfo).Error; err != nil {
			fmt.Println(err)
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (css *caseStepService) EditCaseStep(caseStepInfo *vo.CaseStepInfoVO) interface{} {
	var (
		total float64
	)

	caseStepId := caseStepInfo.StepId
	newCaseStepInfo := &InterfaceTestPartEntity.TItfCaseStepInfo{
		ItfId:    caseStepInfo.ItfId,
		ReqData:  caseStepInfo.ReqData,
		StepName: caseStepInfo.StepName,
		StepDesc: caseStepInfo.StepDesc,
	}
	newAssertInfos := caseStepInfo.AssertInfos
	newVarInfos := caseStepInfo.Variables

	tx := service.DB.Begin()
	// 修改用例步骤信息表
	if err := tx.Model(&InterfaceTestPartEntity.TItfCaseStepInfo{}).Where("step_id = ?", caseStepId).Update(newCaseStepInfo).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}

	for _, assertInfo := range newAssertInfos {
		newAssertInfo := &InterfaceTestPartEntity.TAssertInfo{
			StepId:    caseStepId,
			AssertCol: assertInfo.AssertCol,
			Method:    assertInfo.Method,
			ExpValue:  assertInfo.ExpValue,
		}

		if err := tx.Model(&InterfaceTestPartEntity.TAssertInfo{}).Where("id = ?", assertInfo.ID).Count(&total).Error; err != nil {
			return err
		} else {
			if total == 0 {
				newAssertInfo.AssertId = utils.GenerateUUID()
				if err := tx.Create(newAssertInfo).Error; err != nil {
					fmt.Println(err)
					tx.Rollback()
					return nil
				}
			} else {
				if err := tx.Model(&InterfaceTestPartEntity.TAssertInfo{}).Where("id = ?", assertInfo.ID).Update(newAssertInfo).Error; err != nil {
					tx.Rollback()
					fmt.Println(err)
					return err
				}
			}
		}

	}

	for _, varInfo := range newVarInfos {
		newVarInfo := &InterfaceTestPartEntity.TCaseStepVarInfo{
			StepId:          caseStepId,
			CollectCol:      varInfo.CollectCol,
			CollectColAlias: varInfo.CollectColAlias,
		}

		if err := tx.Model(&InterfaceTestPartEntity.TCaseStepVarInfo{}).Where("id = ?", varInfo.ID).Count(&total).Error; err != nil {
			return err
		} else {
			if total == 0 {
				newVarInfo.VariableId = utils.GenerateUUID()
				if err := tx.Create(newVarInfo).Error; err != nil {
					fmt.Println(err)
					tx.Rollback()
					return nil
				}
			} else {
				if err := tx.Model(&InterfaceTestPartEntity.TCaseStepVarInfo{}).Where("id = ?", varInfo.ID).Update(newVarInfo).Error; err != nil {
					tx.Rollback()
					fmt.Println(err)
					return err
				}
			}
		}

	}

	//// 修改变量表
	//if err := tx.Where("step_id = ?", caseStepId).Update(newVarInfo).Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//// 修改断言表
	//if err := tx.Where("step_id = ?", caseStepId).Update(newAssertInfo).Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}
	tx.Commit()
	return nil
}

func (css *caseStepService) DelCase(json *qjson.QJson) error {
	caseStepId := json.GetString("caseStepId")
	// 目前为物理删除，且只删除了用例步骤表
	if err := service.DB.Where("step_id = ?", caseStepId).Delete(&InterfaceTestPartEntity.TItfCaseStepInfo{}).Error; err != nil {
		return err
	}
	return nil
}

func (css *caseStepService) GetStepList(qj *qjson.QJson) (pageInfo *model.PageInfo, err error) {
	var (
		ret      []*vo.CaseStepInfoVO
		pageNum  = qj.GetNum("pageNum")
		pageSize = qj.GetNum("pageSize")
		caseId   = qj.GetString("caseId")
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

	//hdb := dbUtil.NewHDB(service.DB, &ret)
	//hdb.Paginate(qj).Find(&ret)
	//pageInfo, err = hdb.Pack()
	//midDB := service.DB.Model(&InterfaceTestPartEntity.TItfCaseStepInfo{}).Where("case_id = ?", caseId)
	midDB := service.DB.Table("t_itf_case_step_infos t1").Select("t1.step_id, t1.step_name, IFNULL(t2.step_status, '未运行') as step_status, t2.step_log").Joins(fmt.Sprintf("left join (select t3.step_status, t3.step_log ,t3.step_his_id, t3.step_id, max(t3.created_at) from t_itf_case_step_run_his t3 where t3.case_id = '%s' group by t3.step_id) t2 on t1.step_id = t2.step_id", caseId)).Where("t1.case_id = ? and t1.deleted_at is null", caseId)
	pageInfo, err = dbUtil.Paginate(midDB, pageNum, pageSize, &ret)
	return pageInfo, nil
}

func (css *caseStepService) GetCaseStepDetail(json *qjson.QJson) (*vo.CaseStepInfoVO, error) {
	var (
		stepId      = json.GetString("stepId")
		stepInfo    = InterfaceTestPartEntity.TItfCaseStepInfo{}
		varInfos    = []InterfaceTestPartEntity.TCaseStepVarInfo{}
		assertInfos = []InterfaceTestPartEntity.TAssertInfo{}
	)
	if err := service.DB.Where("step_id = ?", stepId).Find(&stepInfo).Error; err != nil {
		return nil, err
	}
	if err := service.DB.Where("step_id = ?", stepId).Find(&varInfos).Error; err != nil {
		return nil, err
	}
	if err := service.DB.Where("step_id = ?", stepId).Find(&assertInfos).Error; err != nil {
		return nil, err
	}
	return &vo.CaseStepInfoVO{
		CaseId:      stepInfo.CaseId,
		ItfId:       stepInfo.ItfId,
		StepId:      stepInfo.StepId,
		StepNum:     stepInfo.StepNum,
		StepName:    stepInfo.StepName,
		StepDesc:    stepInfo.StepDesc,
		ReqData:     stepInfo.ReqData,
		ExpRes:      stepInfo.ExpRes,
		Variables:   varInfos,
		AssertInfos: assertInfos,
	}, nil
}
