package test

import (
	"fmt"
	"github.com/yancyzhou/Jd_union_sdk/jduion"
	"github.com/yancyzhou/Jd_union_sdk/union"
	"testing"
)

func TestCategory(t *testing.T) {
	var jd jduion.JdSdk
	jd.AppSecret = conf.AppSecret
	jd.JdAppKey = conf.JdAppKey
	var query jduion.CategoryRequest
	query.Req.Grade = 0
	query.Req.ParentId = 0
	jdUnionOpenCategoryGoodsGet := jd.JDUnionOpenCategoryGoodsGet(query)
	fmt.Println("jdUnionOpenCategoryGoodsGet===>", jdUnionOpenCategoryGoodsGet)
}
