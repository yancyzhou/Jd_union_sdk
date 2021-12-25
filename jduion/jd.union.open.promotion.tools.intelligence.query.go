package jduion

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
