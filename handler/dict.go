package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
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
	dicts := database.ListDict("")

	any := make([]*any.Any, 0)
	for _, dict := range dicts {
		temp := model.GetDictPB(dict)
		result, _ := ptypes.MarshalAny(temp)
		any = append(any, result)
	}

	resp.GenerateListResponseSucc(any)
	return nil
}

func (e *Dict) ListChildren(ctx context.Context, req *wrappers.StringValue, resp *libresp.ListResponse) error {
	log.Logf("Received Dict.ListChildren request %s", req)
	dicts := database.ListDict(req.Value)

	any := make([]*any.Any, 0)
	for _, dict := range dicts {
		temp := model.GetDictPB(dict)
		result, _ := ptypes.MarshalAny(temp)
		any = append(any, result)
	}

	resp.GenerateListResponseSucc(any)
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

func (e *Dict) ListCategory(ctx context.Context, req *wrappers.StringValue, resp *libresp.ListResponse) error {
	return nil
}