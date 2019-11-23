package utils

import (
	"fmt"
	"salotto/model"
	"salotto/service"
	"salotto/utils/qjson"
)

func Pagination(ent interface{}, qj *qjson.QJson) (pageInfo *model.PageInfo, err error) {
	var (
		//ret      []*InterfaceTestPartEntity.InterfaceInfo
		pageNum  = qj.GetInt("pageNum")
		pageSize = qj.GetInt("pageSize")
		total    float64
	)

	if err = service.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(ent).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 获取总条数
	service.DB.Model(ent).Count(&total)

	pageInfo = &model.PageInfo{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
		List:     ent,
	}

	return pageInfo, err
}
