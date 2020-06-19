package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro/go-micro/util/log"
	"github.com/pku-hit/dict/component/database"
	"github.com/pku-hit/dict/model"
	"github.com/pku-hit/dict/proto"

	"github.com/pku-hit/libresp"
)

type Dict struct{}

func (e *Dict) ListRoot(ctx context.Context, req *empty.Empty, resp *libresp.ListResponse) error {
	log.Log("Received Dict.ListRoot request")
	dicts, err := database.ListRoot()
	if err != nil {
		resp.GenerateListResponseWithInfo(libresp.GENERAL_ERROR, err.Error())
		return nil
	}
	resp.GenerateListResponseSucc(model.GetDictsAny(dicts))
	return nil
}

func (e *Dict) ListChildren(ctx context.Context, req *proto.ListChildrenRequest, resp *libresp.ListResponse) error {
	log.Logf("Received Dict.ListChildren request %s", req)
	types := make([]proto.DictType, 0)
	if req.Type != proto.DictType_Node && req.Type != proto.DictType_Group {
		types = append(types, proto.DictType_Node, proto.DictType_Group)
	}
	dicts, err := database.ListChildren(req.ParentId, types...)
	if err != nil {
		resp.GenerateListResponseWithInfo(libresp.GENERAL_ERROR, err.Error())
		return nil
	}

	resp.GenerateListResponseSucc(model.GetDictsAny(dicts))
	log.Info(resp)
	return nil
}

func (e *Dict) AddDict(ctx context.Context, req *proto.AddDictRequest, resp *libresp.GenericResponse) error {
	log.Log("Received Dict.AddDict request")
	dict, err := database.NewDict(req.Category, req.ParentId, req.Code, req.Name, req.PyCode, req.Type, req.Value)
	if err != nil {
		resp.GenerateGenericResponseWithInfo(libresp.GENERAL_ERROR, err.Error())
		return nil
	}
	resp.GenerateGenericResponseSucc(model.GetDictAny(dict))
	return nil
}

func (e *Dict) DelDict(ctx context.Context, req *wrappers.StringValue, resp *libresp.Response) error {
	log.Log("Received Dict.DelDict request")
	err := database.DeleteDict(req.Value, true)
	if err != nil {
		resp.GenerateResponseWithInfo(libresp.GENERAL_ERROR, err.Error())
	} else {
		resp.GenerateResponseSucc()
	}
	return nil
}

func (e *Dict) GetValue(ctx context.Context, req *wrappers.StringValue, resp *libresp.GenericResponse) error {
	log.Log("Received Dict.GetValue request")
	dict, err := database.ExistDictWithId(req.Value)
	if err != nil {
		resp.GenerateGenericResponseWithInfo(libresp.GENERAL_ERROR, err.Error())
	} else {
		resp.GenerateGenericResponseSucc(model.GetDictAny(dict))
	}
	return nil
}

func (e *Dict) ListCategory(ctx context.Context, req *wrappers.StringValue, resp *libresp.ListResponse) error {
	return nil
}
