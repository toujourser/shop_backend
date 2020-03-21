package services

import "vue_shop/repositories"

type ReportsServices struct {
	repo *repositories.ReportsRepositories
}

func NewReportsServices() *ReportsServices {
	return &ReportsServices{repo: repositories.NewReportsRepositories()}
}

func (s *ReportsServices)Reports()(map[string]interface{}, error){
	return s.repo.Reports()
}
