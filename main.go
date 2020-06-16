package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/pku-hit/dict/handler"
	_ "github.com/pku-hit/dict/model/entity"
	"github.com/pku-hit/dict/proto"
	"github.com/pku-hit/dict/subscriber"
)

func main() {

	// New Service
	service := micro.NewService(
		micro.Registry(etcdv3.NewRegistry()),
		micro.Name(proto.ServiceName),
		micro.Version(proto.Version),
	)

	// Initialise service
	service.Init()

	// Register Handler
	proto.RegisterDictHandler(service.Server(), new(handler.Dict))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(proto.ServiceName, service.Server(), new(subscriber.Dict))

	// Register Function as Subscriber
	micro.RegisterSubscriber(proto.ServiceName, service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
