package services

import "vue_shop/repositories"

type OrdersServices struct {
	repo *repositories.OrdersRepositories
}

func NewOrdersServices() *OrdersServices {
	return &OrdersServices{repo: repositories.NewOrdersRepositories()}
}

func (s *OrdersServices) List(pageNum, pageSize, user_id int, query, pay_status, is_send, order_fapiao_title, order_fapiao_company, order_fapiao_content, consignee_addr string) (map[string]interface{}, error) {
	return s.repo.Reports(pageNum, pageSize, user_id, query, pay_status, is_send, order_fapiao_title, order_fapiao_company, order_fapiao_content, consignee_addr)
}
