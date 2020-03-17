package services

import (
	"errors"
	"github.com/kataras/iris/v12"
	"vue_shop/repositories"
)

type GoodsServices struct {
	repo *repositories.GoodsReposotories
}

func NewGoodsServices() *GoodsServices {
	return &GoodsServices{repo: repositories.NewGoodsReposotories()}
}

func (s *GoodsServices) List(pageNum, pageSize int, query string) (map[string]interface{}, error) {
	return s.repo.List(pageNum, pageSize, query)
}

func (s *GoodsServices) Create(ctx iris.Context) (map[string]interface{}, error) {
	var goodsObj repositories.GoodsObj
	ctx.ReadJSON(&goodsObj)

	if goodsObj.GoodsName == "" || goodsObj.GoodsPrice == 0 {
		return nil, errors.New("请求参错误!!!")
	}
	return s.repo.Create(goodsObj)
}

func (s *GoodsServices) GetOne(id int) (result map[string]interface{}, err error) {
	return s.repo.GetOne(id)
}

func (s *GoodsServices) Update(id int, ctx iris.Context) (map[string]interface{}, error) {
	var goodsObj repositories.GoodsObj
	ctx.ReadJSON(&goodsObj)

	if goodsObj.GoodsName == "" || goodsObj.GoodsPrice == 0 {
		return nil, errors.New("请求参错误!!!")
	}
	return s.repo.Update(id, goodsObj)
}

func (s *GoodsServices) Delete(id int) error {
	return s.repo.Delete(id)
}
