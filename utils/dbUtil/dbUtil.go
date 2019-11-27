package dbUtil

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"salotto/model"
	"salotto/utils/qjson"
)

type SelfDB struct {
	*gorm.DB
}

func (db *SelfDB) Paginate(ent interface{}, qj *qjson.QJson) (pageInfo *model.PageInfo, err error) {
	var (
		//ret      []*InterfaceTestPartEntity.InterfaceInfo
		pageNum  = qj.GetInt("pageNum")
		pageSize = qj.GetInt("pageSize")
		total    float64
	)

	if err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(ent).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 获取总条数
	db.Model(ent).Count(&total)

	pageInfo = &model.PageInfo{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
		List:     ent,
	}

	return pageInfo, err
}
