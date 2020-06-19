package database

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
	"github.com/pku-hit/dict/model/entity"
	"github.com/pku-hit/dict/proto"
	"github.com/pku-hit/dict/util"
	"time"
)

func DeleteDict(id string, soft bool) (err error) {
	dict, err := ExistDictWithId(id)
	if err == gorm.ErrRecordNotFound {
		return
	}

	now := time.Now()
	if soft {
		dict.Status = proto.DictStatus_Deleted.String()
		dict.DeleteAt = &now
		log.Info(util.Json.StructToMap(dict))
		err = db.Model(dict).Select("status", "delete_at").Update(util.Json.StructToMap(dict)).Error
	} else {
		err = db.Where("id = ?", id).Delete(new(entity.DictInfo)).Error
	}
	return
}

func ExistDictWithId(id string) (dict *entity.DictInfo, err error) {
	dict = &entity.DictInfo{}
	err = db.Where("id = ?", id).Find(dict).Error
	if err == gorm.ErrRecordNotFound {
		dict = nil
	}
	return
}

func ExistDict(parentId, code string) (dict *entity.DictInfo, err error) {
	query := db.New()
	if !util.String.IsEmptyString(parentId) {
		query.Where("ParentId = ?", parentId).Where("Type in (?)", []string{proto.DictType_Group.String(), proto.DictType_Node.String()})
	} else {
		query.Where("ParentId is null").Where("Type = ?", proto.DictType_Root.String())
	}
	query.Where("Code = ?", code)
	dict = &entity.DictInfo{}
	err = query.Find(dict).Error

	if err == gorm.ErrRecordNotFound {
		dict = nil
	}
	return
}

func NewDict(category, parentId, code, name, pyCode string, dictType proto.DictType, value interface{}) (dict *entity.DictInfo, err error) {

	if dict, err = ExistDict(parentId, code); err == nil && dict != nil {
		log.Warnf("exist dict: %s", util.Json.ToJsonString(dict))
		return
	}

	now := time.Now()
	dict = &entity.DictInfo{
		ID:       util.Snowflake.GenId(),
		Code:     code,
		Name:     name,
		PyCode:   pyCode,
		Value:    util.Json.ToJsonString(value),
		Status:   proto.DictStatus_Normal.String(),
		CreateAt: &now,
		UpdateAt: &now,
	}

	if util.String.IsEmptyString(parentId) {
		if dictType != proto.DictType_Root {
			err = errors.New("没有指定ParentId的字典，限制仅允许为Root类型")
			return
		}
		dict.Type = proto.DictType_Root.String()
	} else {
		if dictType != proto.DictType_Node && dictType != proto.DictType_Group {
			err = errors.New("指定ParentId的字典，限制仅允许为Group或Node类型")
			return
		}
		if _, err = ExistDictWithId(parentId); err == gorm.ErrRecordNotFound {
			err = errors.New("指定的父节点不存在")
			return
		}
		dict.ParentId = parentId
		dict.Type = dictType.String()
	}

	if !util.String.IsEmptyString(category) {
		dict.Category = category
	}

	err = db.Save(dict).Error
	if err != nil {
		log.Error("save new dict error %s.", err.Error())
	}
	return
}

func ListRoot(category string) (dicts []*entity.DictInfo, err error) {
	query := db.New()
	if !util.String.IsEmptyString(category) {
		query.Where("category = ?", category)
	}
	query.Where("Type = ? and ParentId is null", proto.DictType_Root.String())
	err = query.Find(&dicts).Error
	if err != nil {
		log.Error(util.Json.ToJsonString(err))
	} else {
		log.Info(util.Json.ToJsonString(dicts))
	}
	return
}

func ListChildren(parentId string, dictType ...proto.DictType) (dicts []*entity.DictInfo, err error) {
	query := db.New()
	if len(dictType) > 0 {
		query.Where("type in (?)", dictType)
	}
	if util.String.IsEmptyString(parentId) {
		err = errors.New("未指定父节点ID")
		return
	}
	err = query.Where("parent_id = ?", parentId).Find(&dicts).Error
	return
}
