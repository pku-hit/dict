package util

import (
	"github.com/bwmarrin/snowflake"
	"math/rand"
)

var node *snowflake.Node

func init() {
	node, _ = snowflake.NewNode(rand.Int63n(64))
}

func GenId() string {
	return node.Generate().String()
}
