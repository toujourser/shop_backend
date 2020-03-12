package services

import (
	"errors"
	"vue_shop/models"
	"vue_shop/repositories"
	"vue_shop/utils"
)

type UserServices struct {
	repo *repositories.UserRepositories
}

func NewAuthServices() *UserServices {
	return &UserServices{repo: repositories.NewAuthRepositories()}
}

func (s *UserServices) UserLogin(m map[string]string) (*models.Manager, error) {
	return s.repo.UserLogin(m)
}

func (s *UserServices) UserList(pagenum, pagesize int, query string) (users []map[string]interface{}, total int) {
	return s.repo.UserList(pagenum, pagesize, query)
}

func (s *UserServices) UserState(uid int, ustate bool) (map[string]interface{}, error) {
	return s.repo.UserState(uid, ustate)
}

func (s *UserServices) Create(m map[string]string) (map[string]interface{}, error) {
	return s.repo.Create(m)
}

func (s *UserServices) GetOne(id int) (map[string]interface{}, error) {
	return s.repo.GetOne(id)
}

func (s *UserServices) PutOne(id int, email, mobile string) (map[string]interface{}, error) {
	if !utils.VerifyEmailFormat(email) {
		return nil, errors.New("邮箱格式不正确！！！")
	} else if !utils.VerifyMobile(mobile) {
		return nil, errors.New("手机格式不正确！！！")
	}
	return s.repo.PutOne(id, email, mobile)
}

func (s *UserServices) DeleteOne(id int) error {
	return s.repo.DeleteOne(id)
}

func (s *UserServices)UserImpower(userId , roleId int)(map[string]interface{}, error){
	return s.repo.UserImpower(userId , roleId)

}