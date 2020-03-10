package repositories

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"vue_shop/models"
)

type MenusRepositories struct {
	db *gorm.DB
}

func NewMenusRepositories() *MenusRepositories {
	return &MenusRepositories{db: models.DB.Mysql}
}

func (r *MenusRepositories) List() ([]map[string]interface{}, error) {
	var menus []*models.Permission
	if err := r.db.Find(&menus).Error; err != nil {
		return nil, errors.New("目录数据查询失败")
	}

	lv1 := []*models.Permission{}
	lv2 := []*models.Permission{}
	lv3 := []*models.Permission{}

	for _, m := range menus {
		if m.PsLevel == "0" {
			lv1 = append(lv1, m)
		} else if m.PsLevel == "1" {
			lv2 = append(lv2, m)
		} else if m.PsLevel == "2" {
			lv3 = append(lv3, m)
		}

	}

	menusMap := []map[string]interface{}{}

	for _, p1 := range lv1 {
		psApi := models.PermissionApi{}
		r.db.Model(&p1).Related(&psApi, "PermissionId")
		println("---------------------")
		fmt.Printf("%+v %+v\n", psApi, p1)

		l1Map := map[string]interface{}{
			"id":       p1.PsId,
			"authName": p1.PsName,
			"path":     psApi.PsApiPath,
			"order":    psApi.PsApiOrder,
			"chidren":  make([]map[string]interface{}, 0),
		}
		menusMap = append(menusMap, l1Map)
	}

	for _, l2 := range lv2 {
		psApi := models.PermissionApi{}
		r.db.Model(&l2).Related(&psApi, "PermissionId")
		for _, mapList := range menusMap {
			if l2.PsPid == cast.ToInt(mapList["id"]) {
				l2Map := map[string]interface{}{
					"id":       l2.PsId,
					"authName": l2.PsName,
					"path":     psApi.PsApiPath,
					"order":    psApi.PsApiOrder,
					"chidren":  make([]map[string]interface{}, 0),
				}
				mapList["chidren"] = append(mapList["chidren"].([]map[string]interface{}), l2Map)
			}
		}
	}

	for _, l3 := range lv3 {
		psApi := models.PermissionApi{}
		r.db.Model(&l3).Related(&psApi, "PermissionId")
		// []map[string]interface{}{}
		for _, mapList := range menusMap {
			for _, val := range mapList["chidren"].([]map[string]interface{}) {
				if l3.PsPid == cast.ToInt(val["id"]) {
					l3Map := map[string]interface{}{
						"id":       l3.PsId,
						"authName": l3.PsName,
						"path":     psApi.PsApiPath,
						"order":    psApi.PsApiOrder,
						//"chidren":  make([]map[string]interface{}, 0),
					}
					val["chidren"] = append(val["chidren"].([]map[string]interface{}), l3Map)
					//mapList["chidren"] = append(mapList["chidren"].([]map[string]interface{}), l3Map)
				}
			}
		}
	}

	//for _, item := range menusMap {
	//	fmt.Printf("%+v \n", item)
	//	println("==================")
	//}

	return menusMap, nil
}

/*
	lv1 := []*models.Permission{}
	lv2 := []*models.Permission{}
	lv3 := []*models.Permission{}

	for _, m := range menus {
		if m.PsLevel == "0" {
			lv1 = append(lv1, m)
		} else if m.PsLevel == "1" {
			lv2 = append(lv2, m)
		} else if m.PsLevel == "2" {
			lv3 = append(lv3, m)
		}
	}

	menusMap := make(map[string][]map[string][]string)
	for _, m1 := range lv1 {
		menusMap[m1.PsName] = []map[string][]string{}
	}

	for _, m2 := range lv2 {
		var mp models.Permission
		r.db.Where("ps_id = ?", m2.PsPid).First(&mp)
		for mmk, mmv := range menusMap {
			if mp.PsName == mmk {
				menusMap[mmk] = append(mmv, map[string][]string{
					m2.PsName: make([]string, 0),
				})
			}

		}
	}

	for _, m3 := range lv3 {
		var mp models.Permission
		r.db.Where("ps_id = ?", m3.PsPid).First(&mp)
		for mmk, mmv := range menusMap {
			for subIndexK, subv := range mmv {
				for sunk, sunListV := range subv {
					if mp.PsName == sunk {
						sunListV = append(sunListV, m3.PsName)
						menusMap[mmk][subIndexK][sunk] = sunListV
					}
				}
			}
		}
	}
	return menusMap, nil
*/
