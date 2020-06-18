package model

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pku-hit/dict/model/entity"
	"github.com/pku-hit/dict/proto"
)

func GetDictPB(input *entity.DictInfo) (output *proto.DictItem) {
	if input == nil {
		return
	}
	output = &proto.DictItem{
		DictUniqueId: input.ID,
		Code:         input.Code,
		Name:         input.Name,
		Value:        input.Value,
	}
	switch input.Type {
	case proto.DictType_Root.String():
		output.Type = proto.DictType_Root
	case proto.DictType_Group.String():
		output.Type = proto.DictType_Group
	case proto.DictType_Node.String():
		output.Type = proto.DictType_Node
	}
	switch input.Status {
	case proto.DictStatus_Deleted.String():
		output.Status = proto.DictStatus_Deleted
	case proto.DictStatus_Normal.String():
		output.Status = proto.DictStatus_Normal
	}
	return
}

func GetDictAny(input *entity.DictInfo) (output *any.Any) {
	pb := GetDictPB(input)
	output, _ = ptypes.MarshalAny(pb)
	return
}

func GetDictsAny(input []*entity.DictInfo) (outputs []*any.Any) {
	outputs = make([]*any.Any, len(input))
	for i, dict := range input {
		outputs[i] = GetDictAny(dict)
	}
	return
}
