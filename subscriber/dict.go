package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	"github.com/golang/protobuf/ptypes/struct"
)

type Dict struct{}

func (e *Dict) Handle(ctx context.Context, msg *structpb.Struct) error {
	log.Log("Handler Received message: ", msg.String())
	return nil
}

func Handler(ctx context.Context, msg *structpb.Struct) error {
	log.Log("Function Received message: ", msg.String())
	return nil
}
