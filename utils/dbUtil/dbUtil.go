package dbUtil

import (
	"github.com/jinzhu/gorm"
	"salotto/model"
)

func Paginate(db *gorm.DB, pageNum, pageSize float64, ent interface{}) (pageInfo *model.PageInfo, err error) {
	var total float64
	db.Count(&total)
	db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Scan(ent)
	pageInfo = &model.PageInfo{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
		List:     ent,
	}
	return pageInfo, nil
}
