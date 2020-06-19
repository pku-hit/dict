package util

import (
	"github.com/bwmarrin/snowflake"
	"math/rand"
)

type snowflakeUtil struct {
	node *snowflake.Node
}

var Snowflake *snowflakeUtil

func init() {
	tempNode, _ := snowflake.NewNode(rand.Int63n(64))
	Snowflake = &snowflakeUtil{node: tempNode}
}

func (u *snowflakeUtil) GenId() string {
	return u.node.Generate().String()
}
