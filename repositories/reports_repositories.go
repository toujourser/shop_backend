package repositories

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"vue_shop/models"
)

type ReportsRepositories struct {
	db *gorm.DB
}

func NewReportsRepositories() *ReportsRepositories {
	return &ReportsRepositories{db: models.DB.Mysql}
}

func (r *ReportsRepositories) Reports() (map[string]interface{}, error) {
	var reports []*models.Report_1
	if err := r.db.Find(&reports).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	sites := make([]string, 0, len(reports))
	temp := map[string]struct{}{}
	for _, report := range reports {
		if _, ok := temp[report.Rp1Area]; !ok {
			temp[report.Rp1Area] = struct{}{}
			sites = append(sites, report.Rp1Area)
		}
	}
	fmt.Printf("%+v\n", sites)
	return nil, nil
}
