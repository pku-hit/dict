package database

import (
	"github.com/micro/go-micro/util/log"
	"github.com/pku-hit/dict/component/util"
	"github.com/pku-hit/dict/model/entity"
)

func SaveDict() {
	dict := &entity.DictInfo{}
	dict.ID = util.GenId()
	dict.Code = "code"
	err := db.Save(dict).Error
	log.Info(err)
}

func ListDict(parentId string) (dicts []*entity.DictInfo) {
	query := db.New()
	if len(parentId) == 0 {
		log.Info("looking for root dicts")
		query.Where("ParentId is null")
	} else {
		query.Where("ParentId = ?", parentId)
	}
	err := query.Find(&dicts).Error
	if err != nil {
		log.Error(util.ToJsonString(err))
	} else {
		log.Info(util.ToJsonString(dicts))
	}
	return
}
