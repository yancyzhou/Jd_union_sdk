package jduion

type JDApi interface {
	/*
	 *	set sign with url param
	 */
	SetSignJointUrlParam(Method string, query string) *JdSdk

	/*
	 * 网站/APP来获取的推广链接，功能同宙斯接口的自定义链接转换、 APP领取代码接口通过商品链接、活动链接获取普通推广链接
	 * 支持传入subunionid参数，可用于区分媒体自身的用户ID，该参数可在订单查询接口返回，需向cps-qxsq@jd.com申请权限。
	 * jd_union_open_promotion_common_get
	 */

	JdUnionOpenPromotionCommonGet(query string) *PromotionCommonResult
	JdUnionOpenOrderRowQuery(query string) *OrderRowQueryResult
	JdUnionOpenOrderBonusQuery(query string) *OrderBonusQueryResult
	JdUnionOpenPromotionBySubUnionIdGet(query string) *SubUnionIdResult
	JdUnionOpenGoodsQuery(query string) *GoodsQueryResult
	/*
		@name jd.union.open.category.goods.get 商品类目查询
		@des 根据商品的父类目id查询子类目id信息，通常用获取各级类目对应关系，以便将推广商品归类。业务参数parentId、grade都输入0可查询所有一级类目ID，之后再用其作为parentId查询其子类目。
		@EliteId 1-好券商品,2-京粉APP-jingdong.超级大卖场,3-小程序-jingdong.好券商品,4-京粉APP-jingdong.主题聚惠1-jingdong.服装运动,5-京粉APP-jingdong.主题聚惠2-jingdong.精选家电,6-京粉APP-jingdong.主题聚惠3-jingdong.超市,7-京粉APP-jingdong.主题聚惠4-jingdong.居家生活,10-9.9专区,11-品牌好货-jingdong.潮流范儿,12-品牌好货-jingdong.精致生活,13-品牌好货-jingdong.数码先锋,14-品牌好货-jingdong.品质家电,15-京仓配送,16-公众号-jingdong.好券商品,17-公众号-jingdong.9.9,18-公众号-jingdong.京东配送
	*/
	JdUnionOpenCategoryGoodsGet(query string) *CategoryResponse

	JdUnionOpenGoodsJingFenQuery(query string) *JFResult
	JdUnionOpenGoodsBigFieldQuery(query string) *JdUnionOpenGoodsBigFieldQueryResult
	JdUnionOpenOrderQuery(query string) *OrderResult //jd_union_open_order_query

}
