package repositories

import (
	"errors"
	"github.com/jinzhu/gorm"
	"vue_shop/models"
)

type RightsRepositories struct {
	db *gorm.DB
}

func NewRightRepositories() *RightsRepositories {
	return &RightsRepositories{db: models.DB.Mysql}
}

func (r *RightsRepositories) RightsList(_type string) ([]map[string]interface{}, error) {

	var prs []*models.Permission
	if err := r.db.Find(&prs).Order("ps_level desc").Error; err != nil {
		return nil, errors.New("权限数据查询失败")
	}

	switch _type {
	case "list":
		{
			ps := []map[string]interface{}{}
			for _, item := range prs {
				psApi := models.PermissionApi{}
				r.db.Model(&item).Related(&psApi, "PermissionId")
				t := map[string]interface{}{
					"id":       item.PsId,
					"authName": item.PsName,
					"level":    item.PsLevel,
					"path":     psApi.PsApiPath,
					"pid":      item.PsPid,
				}
				ps = append(ps, t)
			}
			return ps, nil
		}
	case "tree":
		{
			return GetRightsList(prs, r.db)
		}
	default:
		return nil, errors.New("参数错误")
	}
}
