package jduion

import "encoding/json"

type JdUnionOpenPromotionToolsIntelligenceQueryResponse struct {
	JdUnionOpenPromotionToolsIntelligenceQueryResponse struct {
		QueryResult struct {
			Code string `json:"code"`
			Data struct {
				IntelligenceResp struct {
					ReportContent string `json:"reportContent"`
					Essence       string `json:"essence"`
					StartTime     string `json:"startTime"`
					EndTime       string `json:"endTime"`
					Type          string `json:"type"`
					Cid1List      []int  `json:"cid1List"`
					Status        string `json:"status"`
				} `json:"intelligenceResp"`
			} `json:"data"`
			Message string `json:"message"`
		} `json:"queryResult"`
	} `json:"jd_union_open_promotion_tools_intelligence_query_response"`
}

type JdUnionOpenPromotionToolsIntelligenceQueryRequest struct {
	Req struct {
		Title      string `json:"title"`
		Type       int    `json:"type"`
		Cid1List   []int  `json:"cid_1_list"`
		Status     int    `json:"status"`
		Essence    string `json:"essence"`
		PageIndex  int    `json:"pageIndex"`
		PageSize   int    `json:"pageSize"`
		UnionId    int    `json:"unionid"`
		Pid        string `json:"pid"`
		SubUnionId string `json:"subUnionId"`
		SiteId     int    `json:"siteId"`
		PositionId int    `json:"positionId"`
	} `json:"req"`
}

func (J *JdSdk) JdUnionOpenPromotionToolsIntelligenceQuery(query interface{}) (result *JdUnionOpenPromotionToolsIntelligenceQueryResponse) {
	J.SysParams.Method = "jd.union.open.category.goods.get"
	bodyBytes := J.BodyBytes(query)
	//result := &CategoryResponse{}
	//e := json.Unmarshal([]byte(bodyBytes), &result)
	//if e != nil {
	//	panic(e)
	//}
	//categoryResult = &CateGoryResult{}
	e := json.Unmarshal(bodyBytes, &result)
	if e != nil {
		panic(e)
	}
	return result

}
