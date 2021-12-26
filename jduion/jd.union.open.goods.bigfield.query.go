package jduion

import (
	"encoding/json"
)

/*
 * 接口描述：商品详情查询接口,大字段信息
 */

type BigFieldGoodsReq struct {
	GoodsReq struct {
		SkuIds []int64  `json:"skuIds"`
		Fields []string `json:"fields"`
	} `json:"goodsReq"`
}

type JdUnionOpenGoodsBigFieldQueryResponse struct {
	JdUnionOpenGoodsBigFieldQueryResponse struct {
		QueryResult string `json:"queryResult"`
	} `json:"jd_union_open_goods_bigfield_query_responce"`
}

type JdUnionOpenGoodsBigFieldQueryResult struct {
	Code int `json:"code"`
	Data []struct {
		SkuName           string `json:"skuName"`
		VideoBigFieldInfo struct {
			ProductFeatures     string `json:"productFeatures"`
			Image               string `json:"image"`
			MaterialDescription string `json:"material_Description"`
			Comments            string `json:"comments"`
			BoxContents         string `json:"box_Contents"`
			EditerDesc          string `json:"editerDesc"`
			ContentDesc         string `json:"contentDesc"`
			Manual              string `json:"manual"`
			Catalogue           string `json:"catalogue"`
		} `json:"videoBigFieldInfo"`
		Owner            string `json:"owner"`
		BaseBigFieldInfo struct {
			Wdis     string `json:"wdis"`
			WareQD   string `json:"wareQD"`
			PropCode string `json:"propCode"`
		} `json:"baseBigFieldInfo"`
		MainSkuId        int64 `json:"mainSkuId"`
		BookBigFieldInfo struct {
			ProductFeatures string `json:"productFeatures"`
			Image           string `json:"image"`
			Comments        string `json:"comments"`
			RelatedProducts string `json:"relatedProducts"`
			AuthorDesc      string `json:"authorDesc"`
			BookAbstract    string `json:"bookAbstract"`
			EditerDesc      string `json:"editerDesc"`
			ContentDesc     string `json:"contentDesc"`
			Catalogue       string `json:"catalogue"`
			Introduction    string `json:"introduction"`
		} `json:"bookBigFieldInfo"`
		ProductId    int64  `json:"productId"`
		SkuStatus    int    `json:"skuStatus"`
		DetailImages string `json:"detailImages"`
		ImageInfo    struct {
			ImageList []struct {
				Url string `json:"url"`
			} `json:"imageList"`
		} `json:"imageInfo"`
		SkuId        int64 `json:"skuId"`
		CategoryInfo struct {
			Cid1Name string `json:"cid1Name"`
			Cid2Name string `json:"cid2Name"`
			Cid2     int64  `json:"cid2"`
			Cid3Name string `json:"cid3Name"`
			Cid3     int64  `json:"cid3"`
			Cid1     int64  `json:"cid1"`
		} `json:"categoryInfo"`
	} `json:"data"`
	Message string `json:"message"`
}

func (J *JdSdk) JdUnionOpenGoodsBigFieldQuery(query interface{}) (result *JdUnionOpenGoodsBigFieldQueryResult) {
	J.SysParams.Method = "jd.union.open.goods.bigfield.query"
	bodyBytes := J.BodyBytes(query)
	var res JdUnionOpenGoodsBigFieldQueryResponse
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(res.JdUnionOpenGoodsBigFieldQueryResponse.QueryResult), &result); err != nil {
		panic(err)
	}
	return result

}
