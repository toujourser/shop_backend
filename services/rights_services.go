package services

import "vue_shop/repositories"

type RightsServices struct {
	repo *repositories.RightsRepositories
}

func NewRightsServices() *RightsServices {
	return &RightsServices{repo: repositories.NewRightRepositories()}
}

func (this *RightsServices) RightList(_type string) ([]map[string]interface{}, error) {
	return this.repo.RightsList(_type)
}
