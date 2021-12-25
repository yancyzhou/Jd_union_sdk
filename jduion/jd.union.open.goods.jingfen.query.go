package jduion

import (
	"encoding/json"
)

/*
	@name jd.union.open.category.goods.get 商品类目查询
	@des 根据商品的父类目id查询子类目id信息，通常用获取各级类目对应关系，以便将推广商品归类。业务参数parentId、grade都输入0可查询所有一级类目ID，之后再用其作为parentId查询其子类目。
	@EliteId 1-好券商品,2-京粉APP-jingdong.超级大卖场,3-小程序-jingdong.好券商品,4-京粉APP-jingdong.主题聚惠1-jingdong.服装运动,5-京粉APP-jingdong.主题聚惠2-jingdong.精选家电,6-京粉APP-jingdong.主题聚惠3-jingdong.超市,7-京粉APP-jingdong.主题聚惠4-jingdong.居家生活,10-9.9专区,11-品牌好货-jingdong.潮流范儿,12-品牌好货-jingdong.精致生活,13-品牌好货-jingdong.数码先锋,14-品牌好货-jingdong.品质家电,15-京仓配送,16-公众号-jingdong.好券商品,17-公众号-jingdong.9.9,18-公众号-jingdong.京东配送
*/

type ParamJFReq struct {
	GoodsReq GoodsReq `json:"goodsReq"`
}

type GoodsReq struct {
	EliteId      int    `json:"eliteId"`             //1-好券商品,2-京粉APP-jingdong.超级大卖场,3-小程序-jingdong.好券商品,4-京粉APP-jingdong.主题聚惠1-jingdong.服装运动,5-京粉APP-jingdong.主题聚惠2-jingdong.精选家电,6-京粉APP-jingdong.主题聚惠3-jingdong.超市,7-京粉APP-jingdong.主题聚惠4-jingdong.居家生活,10-9.9专区,11-品牌好货-jingdong.潮流范儿,12-品牌好货-jingdong.精致生活,13-品牌好货-jingdong.数码先锋,14-品牌好货-jingdong.品质家电,15-京仓配送,16-公众号-jingdong.好券商品,17-公众号-jingdong.9.9,18-公众号-jingdong.京东配送
	PageIndex    int    `json:"pageIndex,omitempty"` //页码
	PageSize     int    `json:"pageSize,omitempty"`  //每页数量
	SortName     string `json:"sortName,omitempty"`  //排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	Sort         string `json:"sort,omitempty"`      //asc,desc升降序,默认降序
	Pid          string `json:"pid,omitempty"`
	Fields       string `json:"fields,omitempty"`
	GroupId      int    `json:"groupId,omitempty"`
	OwnerUnionId int    `json:"ownerUnionId,omitempty"`
	ForbidTypes  string `json:"forbidTypes,omitempty"`
}

type JdUnionOpenGoodsJingfenQueryResponse struct {
	JdUnionOpenGoodsJingfenQueryResponse struct {
		Result string `json:"queryResult"`
		Code   string `json:"code"`
	} `json:"jd_union_open_goods_jingfen_query_responce"`
}

//type JFCoupon struct {
//	BindType     int     `json:"bindType"`
//	Discount     float64 `json:"discount"`
//	GetEndTime   int64   `json:"getEndTime"`
//	GetStartTime int64   `json:"getStartTime"`
//	Link         string  `json:"link"`
//	PlatformType int     `json:"platformType"`
//	Quota        float64 `json:"quota"`
//	UseEndTime   int64   `json:"useEndTime"`
//	UseStartTime int64   `json:"useStartTime"`
//	IsBest       int     `json:"isBest"`
//}
//
//type JFRestult struct {
//	Code int `json:"code"`
//	Data []struct {
//		CategoryInfo struct {
//			Cid1     int    `json:"cid1"`
//			Cid1Name string `json:"cid1Name"`
//			Cid2     int    `json:"cid2"`
//			Cid2Name string `json:"cid2Name"`
//			Cid3     int    `json:"cid3"`
//			Cid3Name string `json:"cid3Name"`
//		} `json:"categoryInfo"`
//		Comments       int `json:"comments"`
//		CommissionInfo struct {
//			Commission             float64 `json:"commission"`
//			CommissionShare        float64 `json:"commissionShare"`
//			CouponCommission       float64 `json:"couponCommission"`
//			PlusCommissionShare    float64 `json:"plusCommissionShare"`
//			SuperSubsidies         float64 `json:"superSubsidies"`         //超级补贴
//			ShareCommission        float64 `json:"shareCommission"`        //分享佣金
//			TallestShareCommission float64 `json:"tallestShareCommission"` //最高分享佣金+超级补贴
//			CurrentShareCommission float64 `json:"currentShareCommission"` //当前分享佣金+超级补贴
//			StartTime              int64   `json:"startTime"`
//			EndTime                int64   `json:"endTime"`
//		} `json:"commissionInfo"`
//		CouponInfo struct {
//			CouponList []JFCoupon `json:"couponList"`
//		} `json:"couponInfo"`
//		Coupon            JFCoupon `json:"coupon"`
//		GoodCommentsShare float64  `json:"goodCommentsShare"`
//		ImageInfo         struct {
//			ThumbImg  string `json:"thumbImg"`
//			ImageList []struct {
//				Url string `json:"url"`
//			} `json:"imageList"`
//		} `json:"imageInfo"`
//		InOrderCount30Days int64  `json:"inOrderCount30Days"`
//		MaterialUrl        string `json:"materialUrl"`
//		PriceInfo          struct {
//			Price             float64 `json:"price"`
//			LowestPrice       float64 `json:"lowestPrice"`
//			LowestPriceType   int     `json:"lowestPriceType"`
//			LowestCouponPrice float64 `json:"lowestCouponPrice"`
//			HistoryPriceDay   float64 `json:"historyPriceDay"`
//		} `json:"priceInfo"`
//		ShopInfo struct {
//			ShopName string `json:"shopName"`
//			ShopId   int64  `json:"shopId"`
//		} `json:"shopInfo"`
//		SkuId      int64  `json:"skuId"`
//		SkuName    string `json:"skuName"`
//		IsHot      int    `json:"isHot"`
//		Spuid      int64  `json:"spuid"`
//		BrandCode  string `json:"brandCode"`
//		BrandName  string `json:"brandName"`
//		Owner      string `json:"owner"`
//		PinGouInfo struct {
//			PingouPrice   float64 `json:"pingouPrice"`
//			PingouTmCount int64   `json:"pingouTmCount"`
//			PingouUrl     string  `json:"pingouUrl"`
//		} `json:"pinGouInfo"`
//		ResourceInfo struct {
//			EliteId   int    `json:"eliteId"`
//			EliteName string `json:"eliteName"`
//		} `json:"resourceInfo"`
//		InOrderCount30DaysSku int64 `json:"inOrderCount30DaysSku"`
//	} `json:"data"`
//	TotalCount int64  `json:"totalCount"`
//	Message    string `json:"message"`
//	RequestId  string `json:"requestId"`
//}

type JFResult struct {
	Code int `json:"code"`
	Data []struct {
		BookInfo struct {
			Isbn string `json:"isbn"`
		} `json:"bookInfo"`
		MaterialUrl  string `json:"materialUrl"`
		DocumentInfo struct {
			Document string `json:"document"`
			Discount string `json:"discount"`
		} `json:"documentInfo"`
		ImageInfo struct {
			WhiteImage string `json:"whiteImage"`
			ImageList  []struct {
				Url string `json:"url"`
			} `json:"imageList"`
		} `json:"imageInfo"`
		PinGouInfo struct {
			PingouEndTime   int64   `json:"pingouEndTime"`
			PingouPrice     float64 `json:"pingouPrice"`
			PingouTmCount   int64   `json:"pingouTmCount"`
			PingouUrl       string  `json:"pingouUrl"`
			PingouStartTime int64   `json:"pingouStartTime"`
		} `json:"pinGouInfo"`
		ForbidTypes  []int `json:"forbidTypes"`
		ResourceInfo struct {
			EliteId   int    `json:"eliteId"`
			EliteName string `json:"eliteName"`
		} `json:"resourceInfo"`
		SkuLabelInfo struct {
			Is7ToReturn    int64  `json:"is7ToReturn"`
			Fxg            string `json:"fxg"`
			FxgServiceList []struct {
				ServiceName string `json:"serviceName"`
			} `json:"fxgServiceList"`
		} `json:"skuLabelInfo"`
		SkuName   string `json:"skuName"`
		PriceInfo struct {
			LowestPrice       float64 `json:"lowestPrice"`
			LowestCouponPrice float64 `json:"lowestCouponPrice"`
			Price             float64 `json:"price"`
			HistoryPriceDay   int64   `json:"historyPriceDay"`
			LowestPriceType   int     `json:"lowestPriceType"`
		} `json:"priceInfo"`
		Spuid          int64 `json:"spuid"`
		CommissionInfo struct {
			IsLock              int     `json:"isLock"`
			CommissionShare     float64 `json:"commissionShare"`
			PlusCommissionShare float64 `json:"plusCommissionShare"`
			Commission          float64 `json:"commission"`
			StartTime           int64   `json:"startTime"`
			CouponCommission    float64 `json:"couponCommission"`
			EndTime             int64   `json:"endTime"`
		} `json:"commissionInfo"`
		SkuId     int64  `json:"skuId"`
		BrandCode string `json:"brandCode"`
		Owner     string `json:"owner"`
		ShopInfo  struct {
			LogisticsLvyueScore           string  `json:"logisticsLvyueScore"`
			ShopLevel                     float64 `json:"shopLevel"`
			UserEvaluateScore             string  `json:"userEvaluateScore"`
			ScoreRankRate                 string  `json:"scoreRankRate"`
			AfterServiceScore             string  `json:"afterServiceScore"`
			ShopName                      string  `json:"shopName"`
			ShopLabel                     string  `json:"shopLabel"`
			AfsFactorScoreRankGrade       string  `json:"afsFactorScoreRankGrade"`
			ShopId                        int64   `json:"shopId"`
			LogisticsFactorScoreRankGrade string  `json:"logisticsFactorScoreRankGrade"`
			CommentFactorScoreRankGrade   string  `json:"commentFactorScoreRankGrade"`
		} `json:"shopInfo"`
		BrandName   string `json:"brandName"`
		Comments    int64  `json:"comments"`
		SeckillInfo struct {
			SeckillOriPrice  string `json:"seckillOriPrice"`
			SeckillPrice     string `json:"seckillPrice"`
			SeckillStartTime string `json:"seckillStartTime"`
			SeckillEndTime   string `json:"seckillEndTime"`
		} `json:"seckillInfo"`
		CouponInfo struct {
			CouponList []struct {
				UseEndTime   int64   `json:"useEndTime"`
				GetEndTime   int64   `json:"getEndTime"`
				UseStartTime int64   `json:"useStartTime"`
				Quota        float64 `json:"quota"`
				BindType     int     `json:"bindType"`
				Link         string  `json:"link"`
				PlatformType int     `json:"platformType"`
				Discount     float64 `json:"discount"`
				GetStartTime int64   `json:"getStartTime"`
				HotValue     int     `json:"hotValue"`
				IsBest       int     `json:"isBest"`
			} `json:"couponList"`
		} `json:"couponInfo"`
		PreSaleInfo struct {
			DepositWorth     string `json:"depositWorth"`
			BalanceEndTime   string `json:"balanceEndTime"`
			ShipTime         string `json:"shipTime"`
			PreSalePayType   string `json:"preSalePayType"`
			CurrentPrice     string `json:"currentPrice"`
			PreSaleStartTime string `json:"preSaleStartTime"`
			BalanceStartTime string `json:"balanceStartTime"`
			PreSaleEndTime   string `json:"preSaleEndTime"`
			PreSaleStatus    string `json:"preSaleStatus"`
			AmountDeposit    string `json:"amountDeposit"`
			DiscountType     string `json:"discountType"`
			Earnest          string `json:"earnest"`
			PreAmountDeposit string `json:"preAmountDeposit"`
		} `json:"preSaleInfo"`
		VideoInfo struct {
			VideoList []struct {
				Duration  string `json:"duration"`
				High      string `json:"high"`
				PlayType  string `json:"playType"`
				VideoType string `json:"videoType"`
				ImageUrl  string `json:"imageUrl"`
				Width     string `json:"width"`
				PlayUrl   string `json:"playUrl"`
			} `json:"videoList"`
		} `json:"videoInfo"`
		SecondPriceInfoList []struct {
			SecondPriceType string `json:"secondPriceType"`
			SecondPrice     string `json:"secondPrice"`
		} `json:"secondPriceInfoList"`
		DeliveryType      int     `json:"deliveryType"`
		GoodCommentsShare float64 `json:"goodCommentsShare"`
		CategoryInfo      struct {
			Cid1Name string `json:"cid1Name"`
			Cid2Name string `json:"cid2Name"`
			Cid2     int64  `json:"cid2"`
			Cid3Name string `json:"cid3Name"`
			Cid3     int64  `json:"cid3"`
			Cid1     int64  `json:"cid1"`
		} `json:"categoryInfo"`
		InOrderCount30DaysSku int64 `json:"inOrderCount30DaysSku"`
		InOrderCount30Days    int64 `json:"inOrderCount30Days"`
		ReserveInfo           struct {
			Price                string `json:"price"`
			PanicBuyingEndTime   string `json:"panicBuyingEndTime"`
			StartTime            string `json:"startTime"`
			EndTime              string `json:"endTime"`
			Type                 string `json:"type"`
			Status               string `json:"status"`
			PanicBuyingStartTime string `json:"panicBuyingStartTime"`
		} `json:"reserveInfo"`
		PromotionLabelInfoList []struct {
			PromotionLabel   string `json:"promotionLabel"`
			LableName        string `json:"lableName"`
			PromotionLableId string `json:"promotionLableId"`
			StartTime        string `json:"startTime"`
			EndTime          string `json:"endTime"`
		} `json:"promotionLabelInfoList"`
		IsHot   int   `json:"isHot"`
		JxFlags []int `json:"jxFlags"`
	} `json:"data"`
	Message    string `json:"message"`
	TotalCount int    `json:"totalCount"`
}

func (J *JdSdk) JdUnionOpenGoodsJingFenQuery(query interface{}) *JFResult {
	J.SysParams.Method = "jd.union.open.goods.jingfen.query"
	bodyBytes := J.BodyBytes(query)
	result := &JdUnionOpenGoodsJingfenQueryResponse{}
	e := json.Unmarshal([]byte(bodyBytes), &result)
	if e != nil {
		panic(e)
	}
	var goodsResult JFResult
	e = json.Unmarshal([]byte(result.JdUnionOpenGoodsJingfenQueryResponse.Result), &goodsResult)
	if e != nil {
		panic(e)
	}
	return &goodsResult
}
