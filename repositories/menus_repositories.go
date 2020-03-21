package repositories

import (
	"errors"
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

	return GetRightsList(menus, r.db)
}


func GetRightIds(db *gorm.DB) ([]map[string]interface{}, error) {
	var menus []*models.Permission
	db.Find(&menus)
	lv1, lv2, lv3 := SplitLevel(menus)

	menusMap := []map[string]interface{}{}

	for _, p1 := range lv1 {
		l1Map := map[string]interface{}{
			"id":       p1.PsId,
			"children": make([]map[string]interface{}, 0),
		}
		menusMap = append(menusMap, l1Map)
	}

	for _, l2 := range lv2 {
		psApi := models.PermissionApi{}
		db.Model(&l2).Related(&psApi, "PermissionId")
		for _, mapList := range menusMap {
			if l2.PsPid == cast.ToInt(mapList["id"]) {
				l2Map := map[string]interface{}{
					"id":       l2.PsId,
					"children": make([]map[string]interface{}, 0),
				}
				mapList["children"] = append(mapList["children"].([]map[string]interface{}), l2Map)
			}
		}
	}

	for _, l3 := range lv3 {
		psApi := models.PermissionApi{}
		db.Model(&l3).Related(&psApi, "PermissionId")
		// []map[string]interface{}{}
		for _, mapList := range menusMap {
			for _, val := range mapList["children"].([]map[string]interface{}) {
				if l3.PsPid == cast.ToInt(val["id"]) {
					l3Map := map[string]interface{}{
						"id": l3.PsId,
					}
					val["children"] = append(val["children"].([]map[string]interface{}), l3Map)
				}
			}
		}
	}
	return menusMap, nil

}

func GetRightsList(menus []*models.Permission, db *gorm.DB) ([]map[string]interface{}, error) {

	lv1, lv2, lv3 := SplitLevel(menus)

	menusMap := []map[string]interface{}{}

	for _, p1 := range lv1 {
		psApi := models.PermissionApi{}
		db.Model(&p1).Related(&psApi, "PermissionId")
		//fmt.Printf("%+v %+v\n", psApi, p1)

		l1Map := map[string]interface{}{
			"id":       p1.PsId,
			"authName": p1.PsName,
			"path":     psApi.PsApiPath,
			"order":    psApi.PsApiOrder,
			"children": make([]map[string]interface{}, 0),
		}
		menusMap = append(menusMap, l1Map)
	}

	for _, l2 := range lv2 {
		psApi := models.PermissionApi{}
		db.Model(&l2).Related(&psApi, "PermissionId")
		for _, mapList := range menusMap {
			if l2.PsPid == cast.ToInt(mapList["id"]) {
				l2Map := map[string]interface{}{
					"id":       l2.PsId,
					"authName": l2.PsName,
					"path":     psApi.PsApiPath,
					"order":    psApi.PsApiOrder,
					"children": make([]map[string]interface{}, 0),
				}
				mapList["children"] = append(mapList["children"].([]map[string]interface{}), l2Map)
			}
		}
	}

	for _, l3 := range lv3 {
		psApi := models.PermissionApi{}
		db.Model(&l3).Related(&psApi, "PermissionId")
		// []map[string]interface{}{}
		for _, mapList := range menusMap {
			for _, val := range mapList["children"].([]map[string]interface{}) {
				if l3.PsPid == cast.ToInt(val["id"]) {
					l3Map := map[string]interface{}{
						"id":       l3.PsId,
						"authName": l3.PsName,
						"path":     psApi.PsApiPath,
						"order":    psApi.PsApiOrder,
						//"children":  make([]map[string]interface{}, 0),
					}
					val["children"] = append(val["children"].([]map[string]interface{}), l3Map)
					//mapList["children"] = append(mapList["children"].([]map[string]interface{}), l3Map)
				}
			}
		}
	}

	return menusMap, nil
}

func SplitLevel(menus []*models.Permission) (lv1 []*models.Permission, lv2 []*models.Permission, lv3 []*models.Permission) {
	lv1 = []*models.Permission{}
	lv2 = []*models.Permission{}
	lv3 = []*models.Permission{}

	for _, m := range menus {
		if m.PsLevel == "0" {
			lv1 = append(lv1, m)
		} else if m.PsLevel == "1" {
			lv2 = append(lv2, m)
		} else if m.PsLevel == "2" {
			lv3 = append(lv3, m)
		}
	}
	return
}

