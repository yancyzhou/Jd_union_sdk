package test

import (
	"fmt"
	"github.com/yancyzhou/Jd_union_sdk/jduion"
	"github.com/yancyzhou/Jd_union_sdk/union"
	"testing"
)

/*
 * 	接口描述：根据商品的父类目id查询子类目id信息，通常用获取各级类目对应关系，以便将推广商品归类。业务参数parentId、grade都输入0可查询所有一级类目ID，之后再用其作为parentId查询其子类目。
 */
func TestCategory(t *testing.T) {
	var jd jduion.JdSdk
	jd.AppSecret = conf.AppSecret
	jd.JdAppKey = conf.JdAppKey
	var query jduion.CategoryRequest
	query.Req.Grade = 0
	query.Req.ParentId = 0
	jdUnionOpenCategoryGoodsGet := jd.JdUnionOpenCategoryGoodsGet(query)
	fmt.Println("jdUnionOpenCategoryGoodsGet===>", jdUnionOpenCategoryGoodsGet)
}
