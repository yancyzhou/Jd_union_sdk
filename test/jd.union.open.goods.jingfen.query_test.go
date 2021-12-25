package test

import (
	"encoding/json"
	"fmt"
	"github.com/yancyzhou/Jd_union_sdk/jduion"
	conf "github.com/yancyzhou/Jd_union_sdk/union"
	"testing"
)

func TestJdUnionOpenGoodsJingFenQuery(t *testing.T) {
	var jd jduion.JdSdk
	jd.AppSecret = conf.AppSecret
	jd.JdAppKey = conf.JdAppKey
	var query jduion.ParamJFReq
	query.GoodsReq.PageSize = 20
	query.GoodsReq.PageIndex = 1
	query.GoodsReq.EliteId = 2
	query.GoodsReq.Fields = "hotWords,similar"
	jdUnionOpenGoodsJingFenQueryResult := jd.JdUnionOpenGoodsJingFenQuery(query)
	for _, v := range jdUnionOpenGoodsJingFenQueryResult.Data {
		marshal, _ := json.Marshal(v)
		fmt.Println(string(marshal))
	}
}
