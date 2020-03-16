package services

import (
	"errors"
	"github.com/spf13/cast"
	"vue_shop/repositories"
)

type CategoriesService struct {
	repo *repositories.CategoriesReposotories
}

func NewCategoriesService() *CategoriesService {
	return &CategoriesService{repo: repositories.NewCategoriesReposotories()}
}

func (s *CategoriesService) List(_type, pageNum, pageSize int) map[string]interface{} {
	return s.repo.List(_type, pageNum, pageSize)
}

func (s *CategoriesService) Create(params map[string]interface{}) (result map[string]interface{}, err error) {
	if params["cat_pid"] == "" || params["cat_name"] == "" || params["cat_level"] == "" {
		return nil, errors.New("参数错误！！！")
	}
	return s.repo.Create(cast.ToInt(params["cat_pid"]), cast.ToInt(params["cat_level"]), cast.ToString(params["cat_name"]))
}

func (s *CategoriesService) GetOne(id int) (map[string]interface{}, error) {
	return s.repo.GetOrUpdate(id, "")
}

func (s *CategoriesService) Update(id int, params map[string]interface{}) (map[string]interface{}, error) {
	if params["cat_name"] == "" {
		return nil, errors.New("分类名不能为空！！！")
	}
	return s.repo.GetOrUpdate(id, cast.ToString(params["cat_name"]))
}

func (s *CategoriesService) DeleteOne(id int) error {
	return s.repo.DeleteOne(id)
}

// // 获取分类的属性列表
func (s *CategoriesService) GetAttributes(id int64, sel string) ([]map[string]interface{}, error) {
	if sel != "only" && sel != "many" {
		return nil, errors.New("参数错误")
	}

	return s.repo.GetAttributes(id, sel)
}

// 添加动态参数或者静态属性
func (s *CategoriesService) CreateAttributes(id int64, params map[string]interface{}) (map[string]interface{}, error) {
	if params["attr_name"] == "" || params["attr_sel"] == "" {
		return nil, errors.New("请求参数错误!!!")
	}
	return s.repo.CreateAttributes(id, cast.ToString(params["attr_name"]), cast.ToString(params["attr_sel"]), cast.ToString(params["attr_vals"]))
}

func (s *CategoriesService) DeleteAttributes(cateId int, attrId int) error {
	return s.repo.DeleteAttributes(cateId, attrId)
}

func (s *CategoriesService) GetAttributesByAttrId(cateId int, attrId int, attr_sel, attr_vals string) (map[string]interface{}, error) {
	return s.repo.GetAttributesByAttrId(cateId, attrId, attr_sel, attr_vals)
}
