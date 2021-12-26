package test

import (
	"fmt"
	"github.com/yancyzhou/Jd_union_sdk/jduion"
	conf "github.com/yancyzhou/Jd_union_sdk/union"
	"testing"
)

func TestBigField(t *testing.T) {
	var jd jduion.JdSdk
	jd.AppSecret = conf.AppSecret
	jd.JdAppKey = conf.JdAppKey
	var query jduion.BigFieldGoodsReq
	//query.GoodsReq.Fields = []string{}
	query.GoodsReq.SkuIds = []int64{10041733514084}
	jdUnionOpenGoodsBigFieldQuery := jd.JdUnionOpenGoodsBigFieldQuery(query)
	//fmt.Println("jdUnionOpenGoodsBigFieldQuery===>", jdUnionOpenGoodsBigFieldQuery.Data)
	for _, datum := range jdUnionOpenGoodsBigFieldQuery.Data {
		fmt.Println(datum.SkuId)
	}
}
