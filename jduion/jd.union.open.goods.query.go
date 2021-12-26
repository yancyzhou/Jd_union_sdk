package jduion

import "encoding/json"

type Coupon struct {
	BindType     int     `json:"bindType"`
	Discount     float64 `json:"discount"`
	GetEndTime   int64   `json:"getEndTime"`
	GetStartTime int64   `json:"getStartTime"`
	Link         string  `json:"link"`
	PlatformType int     `json:"platformType"`
	Quota        float64 `json:"quota"`
	UseEndTime   int64   `json:"useEndTime"`
	UseStartTime int64   `json:"useStartTime"`
	IsBest       int     `json:"isBest"`
}

type GoodsQueryResult struct {
	SimilarSkuList []int  `json:"similarSkuList"`
	HotWords       string `json:"hotWords"`
	Code           int    `json:"code"`
	Data           []struct {
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
		SkuLabelInfo struct {
			Is7ToReturn    int `json:"is7ToReturn"`
			Fxg            int `json:"fxg"`
			FxgServiceList []struct {
				ServiceName string `json:"serviceName"`
			} `json:"fxgServiceList"`
		} `json:"skuLabelInfo"`
		SkuName   string `json:"skuName"`
		PriceInfo struct {
			LowestPrice       float64 `json:"lowestPrice"`
			LowestCouponPrice float64 `json:"lowestCouponPrice"`
			Price             float64 `json:"price"`
			HistoryPriceDay   int     `json:"historyPriceDay"`
			LowestPriceType   int     `json:"lowestPriceType"`
		} `json:"priceInfo"`
		StockState int `json:"stockState"`
		SpecInfo   struct {
			Size     string `json:"size"`
			Color    string `json:"color"`
			SpecName string `json:"specName"`
			Spec     string `json:"spec"`
		} `json:"specInfo"`
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
		IsJdSale  int    `json:"isJdSale"`
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
			SeckillOriPrice  float64 `json:"seckillOriPrice"`
			SeckillPrice     float64 `json:"seckillPrice"`
			SeckillStartTime int64   `json:"seckillStartTime"`
			SeckillEndTime   int64   `json:"seckillEndTime"`
		} `json:"seckillInfo"`
		CouponInfo struct {
			CouponList []struct {
				UseEndTime    int64   `json:"useEndTime"`
				GetEndTime    int64   `json:"getEndTime"`
				IsInputCoupon string  `json:"isInputCoupon"`
				UseStartTime  int64   `json:"useStartTime"`
				Quota         float64 `json:"quota"`
				BindType      int     `json:"bindType"`
				Link          string  `json:"link"`
				PlatformType  int     `json:"platformType"`
				Discount      float64 `json:"discount"`
				GetStartTime  int64   `json:"getStartTime"`
				HotValue      int     `json:"hotValue"`
				IsBest        int     `json:"isBest"`
			} `json:"couponList"`
		} `json:"couponInfo"`
		PreSaleInfo struct {
			DepositWorth     float64 `json:"depositWorth"`
			BalanceEndTime   int64   `json:"balanceEndTime"`
			ShipTime         int64   `json:"shipTime"`
			PreSalePayType   int     `json:"preSalePayType"`
			CurrentPrice     float64 `json:"currentPrice"`
			PreSaleStartTime int64   `json:"preSaleStartTime"`
			BalanceStartTime int64   `json:"balanceStartTime"`
			PreSaleEndTime   int64   `json:"preSaleEndTime"`
			PreSaleStatus    int     `json:"preSaleStatus"`
			AmountDeposit    float64 `json:"amountDeposit"`
			DiscountType     int     `json:"discountType"`
			Earnest          float64 `json:"earnest"`
			PreAmountDeposit float64 `json:"preAmountDeposit"`
		} `json:"preSaleInfo"`
		VideoInfo struct {
			VideoList []struct {
				Duration  int64  `json:"duration"`
				High      int    `json:"high"`
				PlayType  string `json:"playType"`
				VideoType int    `json:"videoType"`
				ImageUrl  string `json:"imageUrl"`
				Width     int    `json:"width"`
				PlayUrl   string `json:"playUrl"`
			} `json:"videoList"`
		} `json:"videoInfo"`
		SecondPriceInfoList []struct {
			SecondPriceType int     `json:"secondPriceType"`
			SecondPrice     float64 `json:"secondPrice"`
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
		InOrderCount30Days int64 `json:"inOrderCount30Days"`
		ReserveInfo        struct {
			Price                float64 `json:"price"`
			PanicBuyingEndTime   int64   `json:"panicBuyingEndTime"`
			StartTime            int64   `json:"startTime"`
			EndTime              int64   `json:"endTime"`
			Type                 int     `json:"type"`
			Status               int     `json:"status"`
			PanicBuyingStartTime int64   `json:"panicBuyingStartTime"`
		} `json:"reserveInfo"`
		CommentInfo struct {
			CommentList []struct {
				ImageList []struct {
					Url string `json:"url"`
				} `json:"imageList"`
				Content string `json:"content"`
			} `json:"commentList"`
		} `json:"commentInfo"`
		EliteType              []int `json:"eliteType"`
		PromotionLabelInfoList []struct {
			PromotionLabel   string `json:"promotionLabel"`
			LableName        string `json:"lableName"`
			PromotionLableId int64  `json:"promotionLableId"`
			StartTime        int64  `json:"startTime"`
			EndTime          int64  `json:"endTime"`
		} `json:"promotionLabelInfoList"`
		IsHot   int   `json:"isHot"`
		JxFlags []int `json:"jxFlags"`
	} `json:"data"`
	Message    string `json:"message"`
	TotalCount int64  `json:"totalCount"`
}

//获取商品列表

type JdUnionOpenGoodsQueryResponse struct {
	JdUnionOpenGoodsQueryResponse struct {
		Result string `json:"queryResult"`
		Code   string `json:"code"`
	} `json:"jd_union_open_goods_query_responce"`
}

//请求参数

type GoodsQueryRequest struct {
	GoodsReqDTO GoodsReqDTO `json:"goodsReqDTO"`
}

type GoodsReqDTO struct {
	Cid1                 int64   `json:"cid1,omitempty"`                 //一级类目id
	Cid2                 int64   `json:"cid2,omitempty"`                 //二级类目id
	Cid3                 int64   `json:"cid3,omitempty"`                 //三级类目id
	PageIndex            int     `json:"pageIndex,omitempty"`            //页码
	PageSize             int     `json:"pageSize,omitempty"`             //每页数量，单页数最大30，默认20
	SkuIds               []int64 `json:"skuIds,omitempty"`               //skuid 集合(一次最多支持查询100个sku)，数组类型开发时记得加[]
	Keyword              string  `json:"keyword,omitempty"`              //关键词，字数同京东商品名称一致，目前未限制
	PriceFrom            float64 `json:"pricefrom,omitempty"`            //商品价格下限
	PriceTo              float64 `json:"priceto,omitempty"`              //商品价格上限
	CommissionShareStart int     `json:"commissionShareStart,omitempty"` //佣金比例区间开始
	CommissionShareEnd   int     `json:"commissionShareEnd,omitempty"`   // 佣金比例区间结束
	Owner                string  `json:"owner,omitempty"`                //商品类型：自营[g]，POP[p]
	SortName             string  `json:"sortName,omitempty"`             //排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金)
	Sort                 string  `json:"sort,omitempty"`                 //asc,desc升降序,默认降序
	IsCoupon             int     `json:"isCoupon,omitempty"`             //是否是优惠券商品，1：有优惠券，0：无优惠券
	IsPG                 int     `json:"isPG,omitempty"`                 //是否是拼购商品，1：拼购商品，0：非拼购商品
	PinGouPriceStart     float64 `json:"pingouPriceStart,omitempty"`     //拼购价格区间开始
	PinGouPriceEnd       float64 `json:"pingouPriceEnd,omitempty"`       //拼购价格区间结束
	IsHot                int     `json:"isHot,omitempty"`                //是否是爆款，1：爆款商品，0：非爆款商品
	BrandCode            string  `json:"brandCode,omitempty"`            //品牌code
	ShopId               int     `json:"shopId,omitempty"`               //店铺Id
	HasContent           int     `json:"hasContent,omitempty"`           //1：查询内容商品；其他值过滤掉此入参条件。
	HasBestCoupon        int     `json:"hasBestCoupon,omitempty"`        //1：查询有最优惠券商品；其他值过滤掉此入参条件。
	Pid                  string  `json:"pid,omitempty"`                  //联盟id_应用iD_推广位id
	Fields               string  `json:"fields,omitempty"`               //支持出参数据筛选，逗号','分隔，目前可用：videoInfo(视频信息),hotWords(热词),similar(相似推荐商品),documentInfo(段子信息),skuLabelInfo（商品标签）,promotionLabelInfo（商品促销标签）,stockState（商品库存）
	JxFlags              []int   `json:"jxFlags,omitempty"`              //京喜商品类型，1京喜、2京喜工厂直供、3京喜优选，入参多个值表示或条件查询
	CouponUrl            string  `json:"couponUrl,omitempty"`            //优惠券链接
	ForbidTypes          string  `json:"forbidTypes,omitempty"`          //过滤商品
	ShopLevelFrom        float64 `json:"shopLevelFrom,omitempty"`        //支持传入0.0、2.5、3.0、3.5、4.0、4.5、4.9，默认为空表示不筛选评分
	Isbn                 string  `json:"isbn,omitempty"`                 //图书编号
	SpuId                int64   `json:"spuId,omitempty"`                //主商品spuId
	DeliveryType         int     `json:"deliveryType,omitempty"`         //京东配送 1：是，0：不是
	EliteType            []int   `json:"eliteType,omitempty"`            //资源位17：极速版商品
	IsSeckill            int     `json:"isSeckill,omitempty"`            //是否秒杀商品。1：是
	IsPresale            int     `json:"isPresale,omitempty"`            //是否预售商品。1：是
	IsReserve            int     `json:"isReserve,omitempty"`            //是否预约商品。1:是
	BonusId              int     `json:"bonusId,omitempty"`              //奖励活动ID
}

//关键词字商品查询

func (J *JdSdk) JdUnionOpenGoodsQuery(query interface{}) (GoodResult *GoodsQueryResult) {
	J.SysParams.Method = "jd.union.open.goods.query"
	bodyBytes := J.BodyBytes(query)
	result := &JdUnionOpenGoodsQueryResponse{}
	e := json.Unmarshal(bodyBytes, &result)
	if e != nil {
		panic(e)
	}
	GoodResult = &GoodsQueryResult{}
	e = json.Unmarshal([]byte(result.JdUnionOpenGoodsQueryResponse.Result), GoodResult)
	if e != nil {
		panic(e)
	}
	return GoodResult
}
