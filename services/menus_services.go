package services

import "vue_shop/repositories"

type MenusServices struct {
	repo *repositories.MenusRepositories
}

func NewMenusServices() *MenusServices {
	return &MenusServices{repo: repositories.NewMenusRepositories()}
}

func (s *MenusServices) List() ([]map[string]interface{}, error) {
	return s.repo.List()
}
