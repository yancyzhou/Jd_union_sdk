package test

import (
	"fmt"
	"github.com/yancyzhou/Jd_union_sdk/jduion"
	conf "github.com/yancyzhou/Jd_union_sdk/union"
	"testing"
)

func TestJdUnionOpenPromotionToolsIntelligenceQuery(t *testing.T) {
	var jd jduion.JdSdk
	jd.AppSecret = conf.AppSecret
	jd.JdAppKey = conf.JdAppKey
	var query jduion.JdUnionOpenPromotionToolsIntelligenceQueryRequest
	query.Req.PageSize = 20
	query.Req.PageIndex = 1
	query.Req.Essence = "0"
	jdUnionOpenCategoryGoodsGet := jd.JdUnionOpenPromotionToolsIntelligenceQuery(query)
	fmt.Println("JdUnionOpenPromotionToolsIntelligenceQuery===>", jdUnionOpenCategoryGoodsGet)
}
