package util

import (
	"github.com/pku-hit/dict/model/entity"
	"testing"
)

func TestNilIfEmpty(t *testing.T) {
	dict := &entity.DictInfo{
		ID:       GenId(),
		Category: NilIfEmpty(""),
	}
	t.Log(ToJsonString(dict))
}

func TestToJsonString(t *testing.T) {
	t.Log(ToJsonString("123"))
}