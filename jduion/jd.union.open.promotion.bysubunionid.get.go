package jduion

import (
	"encoding/json"
	"github.com/golang/glog"
)

//转换链接

type JdUnionOpenPromotionBySubUnionIdGetResponse struct {
	JdUnionOpenPromotionBySubUnionIdGetResponse struct {
		Result string `json:"getResult"`
		Code   string `json:"code"`
	} `json:"jd_union_open_promotion_bysubunionid_get_responce"`
}

type SubUnionIdResult struct {
	Code int `json:"code"`
	Data struct {
		ClickURL string `json:"clickURL"`
		ShortURL string `json:"shortURL"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

//转换链接

func (J *JdSdk) JdUnionOpenPromotionBySubUnionIdGet(Query interface{}) (res *SubUnionIdResult) {
	J.SysParams.Method = "jd.union.open.promotion.bysubunionid.get"
	bodyBytes := J.BodyBytes(Query)
	var result JdUnionOpenPromotionBySubUnionIdGetResponse
	e := json.Unmarshal(bodyBytes, &result)

	if e != nil {
		panic(e)
	}
	res = &SubUnionIdResult{}
	e = json.Unmarshal([]byte(result.JdUnionOpenPromotionBySubUnionIdGetResponse.Result), res)
	if e != nil {
		glog.Fatal(e)
		//panic(e)
	}
	return res
}

//业务参数

type PromotionCodeReq struct {
	MaterialId string `json:"materialId"`           //推广物料链接，建议链接使用微Q前缀，能较好适配微信手Q页面
	SubUnionId string `json:"subUnionId,omitempty"` //子联盟ID（需要联系运营开通权限才能拿到数据）
	PositionId int64  `json:"positionId,omitempty"` //推广位ID
	Pid        string `json:"pid,omitempty"`        //子帐号身份标识，格式为子站长ID_子站长网站ID_子站长推广位ID
	CouponUrl  string `json:"couponUrl,omitempty"`  //优惠券领取链接，在使用优惠券、商品二合一功能时入参，且materialId须为商品详情页链接
	ChainType  int    `json:"chainType,omitempty"`  //转链类型，1：长链， 2 ：短链 ，3： 长链+短链，默认短链
}

type BySubUnionIdRequest struct {
	PromotionCodeReq PromotionCodeReq `json:"promotionCodeReq"`
}
