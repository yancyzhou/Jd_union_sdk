package jduion

import (
	"encoding/json"
)

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
		Title      string `json:"title,omitempty"`
		Type       int    `json:"type,omitempty"`
		Cid1List   []int  `json:"cid_1_list,omitempty"`
		Status     int    `json:"status,omitempty"`
		Essence    string `json:"essence,omitempty"`
		PageIndex  int    `json:"pageIndex,omitempty"`
		PageSize   int    `json:"pageSize,omitempty"`
		UnionId    int    `json:"unionid,omitempty"`
		Pid        string `json:"pid,omitempty"`
		SubUnionId string `json:"subUnionId,omitempty"`
		SiteId     int    `json:"siteId,omitempty"`
		PositionId int    `json:"positionId,omitempty"`
	} `json:"req"`
}

func (J *JdSdk) JdUnionOpenPromotionToolsIntelligenceQuery(query interface{}) (result *JdUnionOpenPromotionToolsIntelligenceQueryResponse) {
	J.SysParams.Method = "jd.union.open.promotion.tools.intelligence.query"
	bodyBytes := J.BodyBytes(query)
	e := json.Unmarshal(bodyBytes, &result)
	if e != nil {
		panic(e)
	}
	return result

}
