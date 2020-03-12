package services

import (
	"errors"
	"vue_shop/repositories"
)

type RolesService struct {
	repo *repositories.RolesRepositories
}

func NewRolesService() *RolesService {
	return &RolesService{repo: repositories.NewRolesRepositories()}
}

func (r *RolesService) List() ([]map[string]interface{}, error) {
	return r.repo.List()
}

func (r *RolesService) Create(name, desc string) (map[string]interface{}, error) {
	if name == "" || desc == "" {
		return nil, errors.New("角色名称或描述不能为空！！！")
	}
	return r.repo.Create(name, desc)
}

func (r *RolesService) GetOne(id int) (result map[string]interface{}, err error) {
	return r.repo.GetOne(id)
}

func (r *RolesService) Update(id int, name, desc string) (map[string]interface{}, error) {
	if name == "" || desc == "" {
		return nil, errors.New("角色名称或描述不能为空！！！")
	}
	return r.repo.Update(id, name, desc)
}

func (r *RolesService) DeleteOne(id int) error {
	return r.repo.DeleteOne(id)
}

func (r *RolesService) DeleteRight(roleId, rightId int) (map[string]interface{}, error) {
	return r.repo.DeleteRight(roleId, rightId)
}

func (r *RolesService) RoleImpower(roleId int, rids string) error {
	return r.repo.RoleImpower(roleId, rids)
}
