package services

import "vue_shop/repositories"

type GoodsServices struct {
	repo *repositories.GoodsReposotories
}

func NewGoodsServices() *GoodsServices {
	return &GoodsServices{repo: repositories.NewGoodsReposotories()}
}
