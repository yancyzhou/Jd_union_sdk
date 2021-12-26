package test

import (
	"encoding/json"
	"fmt"
	"github.com/yancyzhou/Jd_union_sdk/jduion"
	conf "github.com/yancyzhou/Jd_union_sdk/union"
	"testing"
)

func TestGoodsQuery(t *testing.T) {
	var jd jduion.JdSdk
	jd.AppSecret = conf.AppSecret
	jd.JdAppKey = conf.JdAppKey
	var query jduion.GoodsQueryRequest
	query.GoodsReqDTO.PageIndex = 1
	query.GoodsReqDTO.PageSize = 20
	query.GoodsReqDTO.Keyword = "iphone13"
	JdUnionOpenGoodsQueryResult := jd.JdUnionOpenGoodsQuery(query)
	fmt.Println("JdUnionOpenGoodsQueryResult===>", JdUnionOpenGoodsQueryResult)
	for _, v := range JdUnionOpenGoodsQueryResult.Data {
		marshal, _ := json.Marshal(v)
		fmt.Println(string(marshal))
	}
}
