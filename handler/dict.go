package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"github.com/micro/go-micro/util/log"
	"github.com/pku-hit/dict/proto"

	"github.com/pku-hit/libresp"
)

type Dict struct{}

func (e *Dict) ListRoot(ctx context.Context, req *empty.Empty, resp *libresp.Response) error {
	log.Log("Received Dict.ListRoot request")
	resp.GenerateResponseSucc()
	return nil
}

func (e *Dict) ListChildren(ctx context.Context, req *wrappers.StringValue, resp *libresp.ListResponse) error {
	log.Log("Received Dict.ListChildren request")
	value := &proto.DictItem{DictUniqueId:uuid.New().String(), Code:"code", Name:"name", Type:proto.DictType_Root}
	result, _ := ptypes.MarshalAny(value)
	resp.GenerateListResponseSucc([]*any.Any{result})
	log.Info(resp)
	return nil
}

func (e *Dict) AddDict(ctx context.Context, req *proto.AddDictRequest, resp *libresp.Response) error {
	log.Log("Received Dict.AddDict request")
	resp.GenerateResponseSucc()
	return nil
}

func (e *Dict) DelDict(ctx context.Context, req *wrappers.StringValue, resp *libresp.Response) error {
	log.Log("Received Dict.DelDict request")
	resp.GenerateResponseSucc()
	return nil
}

func (e *Dict) GetValue(ctx context.Context, req *wrappers.StringValue, resp *libresp.GenericResponse) error {
	log.Log("Received Dict.GetValue request")
	resp.GenerateGenericResponseSucc(nil)
	return nil
}
