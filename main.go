package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/pku-hit/dict/handler"
	proto "github.com/pku-hit/dict/proto"
	"github.com/pku-hit/dict/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("open.svc.dict"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	proto.RegisterDictHandler(service.Server(), new(handler.Dict))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("open.svc.dict", service.Server(), new(subscriber.Dict))

	// Register Function as Subscriber
	micro.RegisterSubscriber("open.svc.dict", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
