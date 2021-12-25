package jduion

import (
	"encoding/json"
	"github.com/yancyzhou/Jd_union_sdk/common"
)

type CategoryResponse struct {
	JdUnionOpenCategoryGoodsGetResponse struct {
		Result string `json:"getResult"`
		Code   string `json:"code"`
	} `json:"jd_union_open_category_goods_get_responce"`
}

type CateGoryResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Grade    int    `json:"grade"`
		ParentId int    `json:"parentId"`
	} `json:"data"`
	RequestId string `json:"requestId"`
}

//获取京东商品类型列表

type JdSdk struct {
	common.JdSdk
}

func (J *JdSdk) JDUnionOpenCategoryGoodsGet(query interface{}) (categoryResult *CateGoryResult) {
	J.SysParams.Method = "jd.union.open.category.goods.get"
	bodyBytes := J.BodyBytes(query)
	result := &CategoryResponse{}
	e := json.Unmarshal([]byte(bodyBytes), &result)
	if e != nil {
		panic(e)
	}
	categoryResult = &CateGoryResult{}
	e = json.Unmarshal([]byte(result.JdUnionOpenCategoryGoodsGetResponse.Result), categoryResult)
	if e != nil {
		panic(e)
	}
	return categoryResult
}

//请求业务参数

type Req struct {
	ParentId int64 `json:"parentId"` //父类目id(一级父类目为0)
	Grade    int64 `json:"grade"`    //类目级别(类目级别 0，1，2 代表一、二、三级类目)
}
type CategoryRequest struct {
	Req Req `json:"req"`
}
