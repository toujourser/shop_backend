package repositories

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"strings"
	"vue_shop/models"
)

type RolesRepositories struct {
	db *gorm.DB
}

func NewRolesRepositories() *RolesRepositories {
	return &RolesRepositories{db: models.DB.Mysql}
}

func (r *RolesRepositories) List() ([]map[string]interface{}, error) {
	var roles []*models.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, errors.New("角色查询失败！！！")
	}

	result := []map[string]interface{}{}

	for _, role := range roles {
		r1 := map[string]interface{}{
			"id":       role.RoleId,
			"roleName": role.RoleName,
			"roleDesc": role.RoleDesc,
			//"children": []map[string]interface{}{},
		}

		psids := strings.Split(role.PsIds, ",")
		//fmt.Printf("%+v\n", psids[:])
		presList := []*models.Permission{}

		for _, pid := range psids {
			var prs models.Permission
			if err := r.db.First(&prs, cast.ToInt(pid)).Error; err != nil {
				continue
			}
			presList = append(presList, &prs)
		}
		rights, _ := GetRightsList(presList, r.db)
		r1["children"] = rights
		result = append(result, r1)
	}

	return result, nil
}

func (r *RolesRepositories) Create(name, desc string) (map[string]interface{}, error) {
	var role models.Role
	if err := r.db.Where("role_name = ?", name).Find(&role).Error; err == nil {
		return nil, errors.New("角色已存在")
	}
	role.RoleName = name
	role.RoleDesc = desc
	if err := r.db.Create(&role).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	result := map[string]interface{}{
		"roleId":   role.RoleId,
		"roleName": role.RoleName,
		"roleDesc": role.RoleDesc,
	}
	return result, nil
}

func (r *RolesRepositories) GetOne(id int) (result map[string]interface{}, err error) {
	var role models.Role
	if err := r.db.Where("role_id = ?", id).First(&role).Error; err != nil {
		return nil, errors.New("角色信息查询失败！！！")
	}
	result = map[string]interface{}{}
	result["id"] = role.RoleId
	result["roleName"] = role.RoleName
	result["roleDesc"] = role.RoleDesc
	return
}

func (r *RolesRepositories) Update(id int, name, desc string) (map[string]interface{}, error) {
	var role models.Role
	if err := r.db.Where("role_id = ?", id).Find(&role).Error; err != nil {
		return nil, errors.New("角色不存在")
	}
	role.RoleName = name
	role.RoleDesc = desc
	if err := r.db.Model(&role).Update(&role).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	result := map[string]interface{}{
		"roleId":   role.RoleId,
		"roleName": role.RoleName,
		"roleDesc": role.RoleDesc,
	}
	return result, nil
}

func (r *RolesRepositories) DeleteOne(id int) error {
	var role models.Role
	if err := r.db.Where("role_id = ?", id).First(&role).Error; err != nil {
		return errors.New("该角色查询不到！！！")
	}
	if err := r.db.Delete(&role).Error; err != nil {
		return errors.New("角色删除失败！！！")
	}
	return nil
}

func (r *RolesRepositories) DeleteRight(roleId, rightId int) (map[string]interface{}, error) {
	var role models.Role
	var permission models.Permission
	if err := r.db.Where("role_id = ?", roleId).Find(&role).Error; err != nil {
		return nil, errors.New("该角色不存在！！！")
	}

	if err := r.db.Where("ps_id = ?", rightId).Find(&permission).Error; err != nil {
		return nil, errors.New("该权限不存在！！！")
	}

	rightid := cast.ToString(rightId)
	psIds := strings.Split(role.PsIds, ",")

	rightRelMap, _ := GetRightIds(r.db)
	switch cast.ToInt(permission.PsLevel) {
	case 0:
		{
			var ids []int
			for _, right := range rightRelMap {
				if right["id"].(int64) == cast.ToInt64(rightId) {
					for _, l2 := range right["children"].([]map[string]interface{}) {
						ids = append(ids, cast.ToInt(l2["id"].(int64)))
						for _, l3 := range l2["children"].([]map[string]interface{}) {
							ids = append(ids, cast.ToInt(l3["id"].(int64)))
						}
					}
				}
			}
			ids = append(ids, rightId)
			for _, rid := range ids {
				for i := 0; i < len(psIds); i++ {
					if cast.ToInt(psIds[i]) == rid {
						psIds = append(psIds[:i], psIds[i+1:]...)
						i--
					}
				}
			}

			role.PsIds = strings.Join(psIds, ",")
		}
	case 1:
		{
			var ids []int
			for _, right := range rightRelMap {
				for _, l2 := range right["children"].([]map[string]interface{}) {
					if l2["id"].(int64) == cast.ToInt64(rightId) {
						ids = append(ids, cast.ToInt(l2["id"].(int64)))
						for _, l3 := range l2["children"].([]map[string]interface{}) {
							ids = append(ids, cast.ToInt(l3["id"].(int64)))
						}
					}
				}
			}
			for _, rid := range ids {
				for i := 0; i < len(psIds); i++ {
					if cast.ToInt(psIds[i]) == rid {
						psIds = append(psIds[:i], psIds[i+1:]...)
						i--
					}
				}
			}
			role.PsIds = strings.Join(psIds, ",")

		}
	case 2:
		{
			for i := 0; i < len(psIds); i++ {
				if psIds[i] == rightid {
					psIds = append(psIds[:i], psIds[i+1:]...)
					i--
				}
			}
			role.PsIds = strings.Join(psIds, ",")
		}
	}

	r.db.Save(&role)

	r1 := map[string]interface{}{
		"id":       role.RoleId,
		"roleName": role.RoleName,
		"roleDesc": role.RoleDesc,
	}

	var presList []*models.Permission

	for _, pid := range psIds {
		var prs models.Permission
		if err := r.db.First(&prs, cast.ToInt(pid)).Error; err != nil {
			continue
		}
		presList = append(presList, &prs)
	}
	rights, _ := GetRightsList(presList, r.db)
	r1["children"] = rights
	return r1, nil
}

func (r *RolesRepositories) RoleImpower(roleId int, rids string) error {
	var role models.Role
	if err := r.db.Where("role_id = ?", roleId).Find(&role).Error; err != nil {
		return errors.New("该角色不存在！！！")
	}
	role.PsIds = rids
	r.db.Save(&role)
	return nil
}
