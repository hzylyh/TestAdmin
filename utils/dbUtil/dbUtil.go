package dbUtil

import (
	"github.com/jinzhu/gorm"
	"salotto/model"
	"salotto/utils/qjson"
)

type HDB struct {
	DB       *gorm.DB
	Ent      interface{}
	pageNum  float64
	pageSize float64
	total    float64
}

func NewHDB(db *gorm.DB, ent interface{}) *HDB {
	return &HDB{
		DB:  db,
		Ent: ent,
	}
}

func (hdb *HDB) Paginate(qj *qjson.QJson) *gorm.DB {
	hdb.pageNum = qj.GetInt("pageNum")
	hdb.pageSize = qj.GetInt("pageSize")
	hdb.DB.Model(hdb.Ent).Count(&hdb.total)

	return hdb.DB.Limit(hdb.pageSize).Offset((hdb.pageNum - 1) * hdb.pageSize)
}

func (hdb *HDB) Pack() (pageInfo *model.PageInfo, err error) {
	pageInfo = &model.PageInfo{
		PageNum:  hdb.pageNum,
		PageSize: hdb.pageSize,
		Total:    hdb.total,
		List:     hdb.Ent,
	}
	return pageInfo, nil
}
