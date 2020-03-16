package repositories

import (
	"github.com/jinzhu/gorm"
	"vue_shop/models"
)

type GoodsReposotories struct {
	db *gorm.DB
}

func NewGoodsReposotories()*GoodsReposotories{
	return &GoodsReposotories{db:models.DB.Mysql}
}
