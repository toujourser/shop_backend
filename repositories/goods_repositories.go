package repositories

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"strings"
	"time"
	"vue_shop/models"
)

type GoodsReposotories struct {
	db *gorm.DB
}

func NewGoodsReposotories() *GoodsReposotories {
	return &GoodsReposotories{db: models.DB.Mysql}
}

type GoodsObj struct {
	GoodsName      string `json:"goods_name"`
	GoodsCat       string `json:"goods_cat"`
	GoodsPrice     string `json:"goods_price"`
	GoodsNumber    string `json:"goods_number"`
	GoodsWeight    string `json:"goods_weight"`
	GoodsIntroduce string `json:"goods_introduce"`
	Pics           []struct {
		Pic string `json:"pic"`
	} `json:"pics"`
	Attrs []struct {
		AttrId    int    `json:"attr_id"`
		AttrValue string `json:"attr_value"`
	} `json:"attrs"`
}

func (r *GoodsReposotories) List(pageNum, pageSize int, query string) (result map[string]interface{}, err error) {
	var goodsList []*models.Goods
	qs := r.db
	qc := r.db.Model(&models.Goods{})
	s := "%" + query + "%"
	if query != "" {
		qs = qs.Where("goods_name like ? or goods_introduce like ?", s, s)
		qc = qc.Where("goods_name like ? or goods_introduce like ?", s, s)
	}
	var total int
	qc.Count(&total)
	qs.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("goods_id desc").Find(&goodsList)

	var goods []map[string]interface{}
	for _, goodsItem := range goodsList {
		g := resultData(goodsItem)
		goods = append(goods, g)
	}
	result = map[string]interface{}{
		"total":   total,
		"pagenum": pageNum,
		"goods":   goods,
	}
	return
}

func (r *GoodsReposotories) Create(goodsObj GoodsObj) (result map[string]interface{}, err error) {
	var goods models.Goods
	if err := r.db.Where("goods_name = ?", goodsObj.GoodsName).First(&goods).Error; err == nil {
		return nil, errors.New("该商品已存在!!!")
	}

	catIds := strings.Split(goodsObj.GoodsCat, ",")
	goods.GoodsName = goodsObj.GoodsName
	goods.GoodsPrice = cast.ToFloat64(goodsObj.GoodsPrice)
	goods.GoodsNumber = cast.ToInt(goodsObj.GoodsNumber)
	goods.GoodsWeight = cast.ToFloat64(goodsObj.GoodsWeight)
	goods.GoodsIntroduce = goodsObj.GoodsIntroduce
	goods.AddTime = time.Now().Unix()
	goods.UpdTime = time.Now().Unix()

	if len(catIds) > 0 {
		goods.CategoryId = cast.ToInt(catIds[len(catIds)-1])
		goods.CatOneId = cast.ToInt(catIds[0])
		if len(catIds) > 1 {
			goods.CatTwoId = cast.ToInt(catIds[1])
		}
		if len(catIds) > 2 {
			goods.CatThreeID = cast.ToInt(catIds[2])
		}
	}

	if err := r.db.Save(&goods).Error; err != nil {
		return nil, errors.New("商品信息保存失败!!!")
	}

	picsList := []*models.GoodsPics{}
	for _, pic := range goodsObj.Pics {
		var pics models.GoodsPics
		pics.GoodsId = goods.GoodsId
		pics.PicsSma = pic.Pic
		if err := r.db.Save(&pics).Error; err != nil {
			return nil, errors.New("商品图片保存失败!!!")
		}
		picsList = append(picsList, &pics)
	}

	attrsList := []*models.GoodsAttr{}
	for _, att := range goodsObj.Attrs {
		var attrs models.GoodsAttr
		attrs.GoodsId = goods.GoodsId
		attrs.AttributeId = att.AttrId
		attrs.AttrValue = att.AttrValue
		if err := r.db.Save(&attrs).Error; err != nil {
			return nil, errors.New("商品属性信息保存失败!!!")
		}
		attrsList = append(attrsList, &attrs)
	}

	result = resultData(&goods)
	result["pics"] = picsList
	result["attrs"] = attrsList

	return
}

func resultData(goods *models.Goods) (result map[string]interface{}) {
	return map[string]interface{}{
		"goods_id":         goods.GoodsId,
		"goods_name":       goods.GoodsName,
		"goods_price":      goods.GoodsPrice,
		"cat_id":           goods.CategoryId,
		"goods_number":     goods.GoodsNumber,
		"goods_weight":     goods.GoodsWeight,
		"goods_introduce":  goods.GoodsIntroduce,
		"goods_big_logo":   goods.GoodsBigLogo,
		"goods_small_logo": goods.GoodsSmallLogo,
		"goods_state":      goods.GoodsState,
		"add_time":         goods.AddTime,
		"upd_time":         goods.UpdTime,
		"hot_number":       goods.HotNumber,
		"is_promote":       goods.IsPromote,
	}
}

func (r *GoodsReposotories) GetOne(id int) (result map[string]interface{}, err error) {
	var goods models.Goods
	var pics []models.GoodsPics
	var attrs []models.GoodsAttr

	if err := r.db.First(&goods, id).Error; err != nil {
		return nil, errors.New("商品信息查询失败!!!")
	}

	if err := r.db.Where("goods_id = ?", goods.GoodsId).Find(&pics).Error; err != nil {
		return nil, errors.New("商品相册关联信息失败!!!")
	}

	if err := r.db.Where("goods_id = ?", goods.GoodsId).Find(&attrs).Error; err != nil {
		return nil, errors.New("商品属性关联信息失败!!!")
	}

	result = resultData(&goods)
	result["pics"] = pics
	result["attrs"] = attrs
	return
}

func (r *GoodsReposotories) Update(id int, goodsObj GoodsObj) (result map[string]interface{}, err error) {
	var goods models.Goods
	var pics []models.GoodsPics
	var attrs []models.GoodsAttr

	if err := r.db.First(&goods, id).Error; err != nil {
		return nil, errors.New("商品信息不存在!!!")
	}
	if err := r.db.Where("goods_id = ?", goods.GoodsId).Find(&pics).Error; err != nil {
		return nil, errors.New("商品相册关联信息失败!!!")
	}

	if err := r.db.Debug().Where("goods_id = ?", goods.GoodsId).Find(&attrs).Error; err != nil {
		return nil, errors.New("商品属性关联信息失败!!!")
	}

	goods.GoodsName = goodsObj.GoodsName
	goods.GoodsPrice = cast.ToFloat64(goodsObj.GoodsPrice)
	goods.GoodsNumber = cast.ToInt(goodsObj.GoodsNumber)
	goods.GoodsWeight = cast.ToFloat64(goodsObj.GoodsWeight)
	goods.GoodsIntroduce = goodsObj.GoodsIntroduce
	goods.UpdTime = time.Now().Unix()
	pics[0].PicsSma = goodsObj.Pics[0].Pic
	for i, pic := range goodsObj.Pics {
		pics[i].PicsSma = pic.Pic
		r.db.Save(pics[i])
	}

	for i, att := range goodsObj.Attrs {
		attrs[i].AttributeId = att.AttrId
		attrs[i].AttrValue = att.AttrValue
		r.db.Save(attrs[i])
	}

	result = resultData(&goods)
	result["pics"] = pics
	result["attrs"] = attrs
	return
}

func (r *GoodsReposotories) Delete(id int) error {
	var goods models.Goods
	var pics models.GoodsPics
	var attrs models.GoodsAttr
	if err := r.db.First(&goods, id).Error; err != nil {
		return errors.New("商品信息不存在!!!")
	}

	r.db.Delete(&pics, "goods_id = ?", goods.GoodsId)
	r.db.Delete(&attrs, "goods_id = ?", goods.GoodsId)
	r.db.Delete(&goods, goods.GoodsId)
	return nil
}
