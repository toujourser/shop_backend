package repositories

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"time"
	"vue_shop/models"
	"vue_shop/utils"
)

type UserRepositories struct {
	db *gorm.DB
}

func NewAuthRepositories() *UserRepositories {
	return &UserRepositories{db: models.DB.Mysql}
}

func (r *UserRepositories) UserLogin(m map[string]string) (*models.Manager, error) {
	var manager models.Manager
	if err := r.db.Where("mg_name=?", m["username"]).First(&manager).Error; err != nil {
		return nil, errors.New("用户不存在！！！")
	}

	if !utils.CompareAES(manager.MgPwd, m["password"]) {
		return nil, errors.New("用户名或密码错误！！！")
	}
	r.db.Model(&manager).Where("username=?", manager.MgName).Updates(map[string]interface{}{"update_time": time.Now().Unix()})
	return &manager, nil
}

func (r *UserRepositories) UserList(pagenum, pagesize int, query string) (userList []map[string]interface{}, total int) {
	var users = []*models.Manager{}
	qs := r.db
	qc := r.db.Model(&models.Manager{})

	if pagesize == 0 {
		pagesize = 10
	}

	s := "%" + query + "%"
	if query != "" {
		qs = qs.Where("mg_name like ? or mg_email like ?", s, s)
		qc = qc.Where("mg_name like ? or mg_email like ?", s, s)
	}
	qc.Count(&total)
	qs.Limit(pagesize).Offset((pagenum - 1) * pagesize).Order("mg_id desc").Find(&users)

	//userList := []map[string]interface{}{}
	for _, user := range users {
		var role models.Role
		r.db.Model(&role).Where("role_id = ?", user.RoleId).Find(&role)
		u := map[string]interface{}{
			"id":          user.MgId,
			"username":    user.MgName,
			"email":       user.MgEmail,
			"mobile":      user.MgMobile,
			"create_time": user.MgTime,
			"mg_status":   cast.ToBool(user.MgState),
			"role_name":   role.RoleName,
		}
		userList = append(userList, u)
	}

	return
}

func (r *UserRepositories) UserState(uId int, uState bool) (map[string]interface{}, error) {
	var manager models.Manager
	s := 0
	if uState {
		s = 1
	} else {
		s = 0
	}

	if err := r.db.Model(&manager).Where("mg_id = ?", uId).First(&manager).Update("mg_state", s).Error; err != nil {
		return nil, errors.New("用户权限修改失败！！")
	}

	u := map[string]interface{}{
		"id":          manager.MgId,
		"username":    manager.MgName,
		"email":       manager.MgEmail,
		"mobile":      manager.MgMobile,
		"create_time": manager.MgTime,
		"mg_status":   cast.ToBool(manager.MgState),
		//"role_name":   role.RoleName,
	}
	return u, nil
}

func (r *UserRepositories) Create(m map[string]string) (map[string]interface{}, error) {
	var manger models.Manager
	err := r.db.Where("mg_name = ?", m["username"]).First(&manger).Error

	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	u := models.Manager{
		MgName:   m["username"],
		MgPwd:    utils.EncryptAES([]byte(m["password"])),
		MgTime:   cast.ToInt(time.Now().Unix()),
		MgMobile: m["mobile"],
		MgEmail:  m["email"],
		MgState:  1,
	}

	if err := r.db.Create(&u).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	result := map[string]interface{}{
		"id":       u.MgId,
		"username": u.MgName,
		"mobile":   u.MgMobile,
		"email":    u.MgEmail,
	}
	return result, nil
}

func (r *UserRepositories) GetOne(id int) (map[string]interface{}, error) {
	var manager models.Manager
	if err := r.db.Where("mg_id = ?", id).First(&manager).Error; err != nil {
		return nil, errors.New("用户查询失败！！！")
	}

	m := map[string]interface{}{
		"id":       manager.MgId,
		"username": manager.MgName,
		"role_id":  manager.RoleId,
		"mobile":   manager.MgMobile,
		"email":    manager.MgEmail,
	}
	return m, nil
}

func (r *UserRepositories) PutOne(id int, email, mobile string) (map[string]interface{}, error) {
	var manager models.Manager
	if err := r.db.Where("mg_id = ?", id).First(&manager).Error; err != nil {
		return nil, errors.New("用户查询失败！！！")
	}
	manager.MgEmail = email
	manager.MgMobile = mobile
	if err := r.db.Model(&manager).Update(&manager).Error; err != nil {
		return nil, errors.New("用户更新失败！！！")
	}
	m := map[string]interface{}{
		"id":       manager.MgId,
		"username": manager.MgName,
		"role_id":  manager.RoleId,
		"mobile":   manager.MgMobile,
		"email":    manager.MgEmail,
	}
	return m, nil
}

func (r *UserRepositories) DeleteOne(id int) error {
	var manager models.Manager
	if err := r.db.Where("mg_id = ?", id).First(&manager).Error; err != nil {
		return errors.New("该用户查询不到！！！")
	}
	if err := r.db.Delete(&manager).Error; err != nil {
		return errors.New("用户删除失败！！！")
	}
	return nil
}

func (r *UserRepositories) UserImpower(userId, roleId int) (map[string]interface{}, error) {
	var manager models.Manager
	if err := r.db.Where("mg_id = ?", userId).First(&manager).Error; err != nil {
		return nil, errors.New("该用户查询不到！！！")
	}
	var role models.Role
	if err := r.db.Where("role_id = ?", roleId).First(&role).Error; err != nil {
		return nil, errors.New("该角色查询不到！！！")
	}

	manager.RoleId = cast.ToInt64(roleId)
	r.db.Save(&manager)
	roleInfo := map[string]interface{}{
		"id":      manager.MgId,
		"role_id": manager.RoleId,
		"mobile":  manager.MgMobile,
		"email":   manager.MgEmail,
	}
	return roleInfo, nil
}
