package repositories

import (
	"github.com/jinzhu/gorm"
	"vue_shop/models"
)

type UploadFileRepositories struct {
	db *gorm.DB
}

func NewUploadFileRepositories()*UploadFileRepositories{
	return &UploadFileRepositories{db:models.DB.Mysql}
}
