package handler

import (
	"github.com/pku-hit/dict/proto"
	"google.golang.org/grpc"
)

var client *proto.DictClient

func init() {
	
	client = proto.NewDictClient()
}
