package entity

import "time"

type DictInfo struct {
	ID       string `gorm:"primary_key;type:varchar(32)"`
	Type     string `gorm:"column:type;type:varchar(32)"`
	Category string `gorm:"category;type:varchar(64)"`

	Parent   *DictInfo   `gorm:"ForeignKey:ParentId"`
	Children []*DictInfo `gorm:"ForeignKey:ParentId"`
	ParentId string      `gorm:"parent_id;type:varchar(32)"`

	Code     string     `gorm:"code";type:varchar(32)`
	PyCode   string     `gorm:"pycode";type:varchar(32)`
	Name     string     `gorm:"name";type:varchar(32)`
	Value    string     `gorm:"value";type:varchar(255)`
	Status   string     `gorm:"status";type:varchar(32)`
	Remark   string     `gorm:"remark";type:varchar(128)`
	CreateAt *time.Time `gorm:"create_at"`
	UpdateAt *time.Time `gorm:"update_at"`
	DeleteAt *time.Time `gorm:"delete_at"`
}

func (DictInfo) TableName() string {
	return "dict.dict_Info"
}
