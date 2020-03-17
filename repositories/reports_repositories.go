package repositories

import (
	"github.com/jinzhu/gorm"
	"vue_shop/models"
)

type ReportsRepositories struct {
	db *gorm.DB
}

func NewReportsRepositories() *ReportsRepositories {
	return &ReportsRepositories{db: models.DB.Mysql}
}
