package test

import (
	"fmt"
	"github.com/yancyzhou/Jd_union_sdk/jduion"
	conf "github.com/yancyzhou/Jd_union_sdk/union"
	"testing"
)

func TestJdUnionOpenPromotionBySubUnionIdGet(t *testing.T) {
	var jd jduion.JdSdk
	jd.AppSecret = conf.AppSecret
	jd.JdAppKey = conf.JdAppKey
	var query jduion.BySubUnionIdRequest
	query.PromotionCodeReq.MaterialId = "https://item.m.jd.com/product/100010793716.html"
	query.PromotionCodeReq.SubUnionId = "vincent"
	jdUnionOpenPromotionBySubUnionIdGet := jd.JdUnionOpenPromotionBySubUnionIdGet(query)
	fmt.Println("jdUnionOpenPromotionBySubUnionIdGet===>", jdUnionOpenPromotionBySubUnionIdGet)
}
