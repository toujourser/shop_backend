package repositories

import (
	"github.com/jinzhu/gorm"
	"vue_shop/models"
)

type OrdersRepositories struct {
	db *gorm.DB
}

func NewOrdersRepositories() *OrdersRepositories {
	return &OrdersRepositories{db: models.DB.Mysql}
}

func (r *OrdersRepositories) Reports(pageNum, pageSize, user_id int, query, pay_status, is_send, order_fapiao_title, order_fapiao_company, order_fapiao_content, consignee_addr string) (map[string]interface{}, error) {
	var orders []*models.Order
	qs := r.db
	qc := r.db.Model(&models.Order{})
	s := "%" + query + "%"
	if query != "" {
		qs = qs.Where("goods_name like ? or goods_introduce like ?", s, s)
		qc = qc.Where("goods_name like ? or goods_introduce like ?", s, s)
	}
	var total int
	qc.Count(&total)
	qs.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("goods_id desc").Find(&orders)

	return nil, nil
}
