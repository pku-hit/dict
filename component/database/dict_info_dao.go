package database

import (
	"github.com/micro/go-micro/util/log"
	"github.com/pku-hit/dict/component/util"
	"github.com/pku-hit/dict/model/entity"
)

func SaveDict() {
	dict := &entity.DictInfo{}
	dict.ID = util.GenId()
	dict.Type = "type"
	dict.Code = "code"
	err := db.Save(dict).Error
	log.Info(err)
}
