package models

// 订单表
type Order struct {
	OrderId            int64   `gorm:"primary_key;AUTO_INCREMENT" json:"order_id"`
	UserId             int64   `json:"user_id"`
	OrderNumber        string  `gorm:"size:32" json:"order_number"`                    // 订单编号
	OrderPrice         float64 `gorm:"default:0.0" json:"order_price"`                 // 订单总金额
	OrderPay           int     `gorm:"default:1" json:"order_pay"`                     // 支付方式  0未支付 1支付宝  2微信  3银行卡
	IsSend             string  `gorm:"default:'否'" json:"is_send"`                     // 订单是否已经发货
	TradeNo            string  `gorm:"size:32;default:''" json:"trade_no"`             // 支付宝交易流水号码
	OrderFapiaoTitle   string  `gorm:"default:'个人'" json:"order_fapiao_title"`         // 发票抬头 个人 公司
	OrderFapiaoCompany string  `gorm:"size:32;default:''" json:"order_fapiao_company"` // 公司名称
	OrderFapiaoContent string  `gorm:"size:32;default:''" json:"order_fapiao_content"` // 发票内容
	ConsigneeAddr      string  `gorm:"type:text" json:"consignee_addr"`                // consignee收货人地址
	PayStatus          int     `gorm:"default:0" json:"pay_status"`                    // 订单状态： 0未付款、1已付款
	CreateTime         int     `gorm:"size:11" json:"create_time"`
	UpdateTime         int     `gorm:"size:11" json:"update_time"`
}

// 商品订单关联表
type OrderGoods struct {
	Id              int64   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	OrderId         int64   `json:"order_id"`
	GoodsId         int64   `json:"goods_id"`
	GoodsPrice      float64 `json:"goods_price"`
	GoodsNumber     int64   `gorm:"default:1" json:"goods_number"`         // 购买单个商品数量
	GoodsTotalPrice float64 `gorm:"default:0.00" json:"goods_total_price"` // 商品小计价格
}
