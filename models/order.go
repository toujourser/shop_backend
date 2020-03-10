package models

// 订单表
type Order struct {
	OrderId            int64   `orm:"pk" json:"order_id"`
	User               *User   `orm:"rel(fk)" json:"user_id"`
	OrderNumber        string  `orm:"size(32)" json:"order_number"`                     // 订单编号
	OrderPrice         float64 `orm:"default(0.00)" json:"order_price"`                 // 订单总金额
	OrderPay           int     `orm:"default(1)" json:"order_pay"`                      // 支付方式  0未支付 1支付宝  2微信  3银行卡
	IsSend             string  `orm:"default(否)" json:"is_send"`                        // 订单是否已经发货
	TradeNo            string  `orm:"size(32);default('')" json:"trade_no"`             // 支付宝交易流水号码
	OrderFapiaoTitle   string  `orm:"default(个人)" json:"order_fapiao_title"`            // 发票抬头 个人 公司
	OrderFapiaoCompany string  `orm:"size(32);default('')" json:"order_fapiao_company"` // 公司名称
	OrderFapiaoContent string  `orm:"size(32);default('')" json:"order_fapiao_content"` // 发票内容
	ConsigneeAddr      string  `orm:"type(text)" json:"consignee_addr"`                 // consignee收货人地址
	PayStatus          int     `orm:"default(0)" json:"pay_status"`                     // 订单状态： 0未付款、1已付款
	CreateTime         int     `orm:"size(11)" json:"create_time"`
	UpdateTime         int     `orm:"size(11)" json:"update_time"`
}

// 商品订单关联表
type OrderGoods struct {
	Id              int64   `orm:"pk" json:"id"`
	Order           *Order  `orm:"rel(fk)" json:"order_id"`
	Goods           *Goods  `orm:"rel(fk)" json:"goods_id"`
	GoodsPrice      float64 `json:"goods_price"`
	GoodsNumber     int64   `orm:"default(1)" json:"goods_number"`         // 购买单个商品数量
	GoodsTotalPrice float64 `orm:"default(0.00)" json:"goods_total_price"` // 商品小计价格
}
