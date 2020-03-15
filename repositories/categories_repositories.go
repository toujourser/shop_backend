package repositories

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"vue_shop/models"
)

type CategoriesReposotories struct {
	db *gorm.DB
}

func NewCategoriesReposotories() *CategoriesReposotories {
	return &CategoriesReposotories{db: models.DB.Mysql}
}

func (r *CategoriesReposotories) GetOrUpdate(id int, cateName string) (result map[string]interface{}, err error) {
	var cate models.Category
	if err := r.db.First(&cate, id).Error; err != nil {
		return nil, errors.New("该分类查询不到！！！")
	}
	if cateName != "" {
		cate.CatName = cateName
		r.db.Save(&cate)
	}
	result = map[string]interface{}{
		"cat_id":    cate.CatId,
		"cat_name":  cate.CatName,
		"cat_pid":   cate.CatPid,
		"cat_level": cate.CatLevel,
	}
	return
}

func (r *CategoriesReposotories) DeleteOne(id int) error {
	var cate models.Category
	if err := r.db.First(&cate, id).Error; err != nil {
		return errors.New("该分类查询不到！！！")
	}
	cate.CatDeleted = 1
	r.db.Save(&cate)
	return nil
}

func (r *CategoriesReposotories) List(_type, pageNum, pageSize int) map[string]interface{} {
	var data []map[string]interface{}
	switch _type {
	case 1:
		data = getLevel1(r.db, pageNum, pageSize)
	case 2:
		data = getLevel2(r.db, pageNum, pageSize, 2)
	case 3:
		data = getLevel3(r.db, pageNum, pageSize)
	default:
		data = getLevel3(r.db, pageNum, pageSize)
	}
	var total int
	r.db.Model(&models.Category{}).Where("cat_level = ?", 0).Count(&total)
	result := map[string]interface{}{
		"total":    total,
		"pagenum":  pageNum,
		"pagesize": pageSize,
		"result":   data,
	}
	return result

}

// 将categories 结构体转换为map对象
func ParseCataItem2Map(cat *models.Category, level int, l2 ...int) map[string]interface{} {
	c := map[string]interface{}{
		"cat_id":      cat.CatId,
		"cat_name":    cat.CatName,
		"cat_pid":     cat.CatPid,
		"cat_level":   cat.CatLevel,
		"cat_deleted": cast.ToBool(cat.CatDeleted),
	}
	if level == 0 || level == 1 {
		if len(l2) == 0 {
			c["children"] = []map[string]interface{}{}
		}
	}

	return c

}

// 获取第一层分类列表
func getLevel1(db *gorm.DB, pagenum, pagesize int) (result []map[string]interface{}) {
	var lv1 []*models.Category
	db.Debug().Where("cat_level = ?", 0).Limit(pagesize).Offset((pagenum - 1) * pagesize).Find(&lv1)

	for _, cat := range lv1 {
		c := ParseCataItem2Map(cat, 0)
		result = append(result, c)
	}
	return
}

// 获取第二层分类列表
func getLevel2(db *gorm.DB, pagenum, pagesize int, level ...int) []map[string]interface{} {
	var lv2 []*models.Category
	db.Debug().Where("cat_level = ?", 1).Find(&lv2)

	lv1 := getLevel1(db, pagenum, pagesize)

	for _, l2 := range lv2 {
		for _, l1 := range lv1 {
			if cast.ToInt64(l2.CatPid) == l1["cat_id"].(int64) {
				var c map[string]interface{}
				if len(level) != 0 {
					c = ParseCataItem2Map(l2, 1, level[0])
				} else {
					c = ParseCataItem2Map(l2, 1)
				}

				l1["children"] = append(l1["children"].([]map[string]interface{}), c)
			}
		}
	}

	return lv1
}

// 获取第三层分类列表
func getLevel3(db *gorm.DB, pagenum, pagesize int) []map[string]interface{} {
	var lv3 []*models.Category
	db.Debug().Where("cat_level = ?", 2).Find(&lv3)

	lv1Andlv2 := getLevel2(db, pagenum, pagesize)

	for _, l3 := range lv3 {
		for _, l1 := range lv1Andlv2 {
			for _, l2 := range l1["children"].([]map[string]interface{}) {
				if cast.ToInt64(l3.CatPid) == l2["cat_id"].(int64) {
					c := ParseCataItem2Map(l3, 2)
					l2["children"] = append(l2["children"].([]map[string]interface{}), c)
				}
			}
		}
	}
	return lv1Andlv2
}

func (r *CategoriesReposotories) Create(pid, level int, name string) (result map[string]interface{}, err error) {

	if pid != 0 {
		if err = r.db.First(&models.Category{}, pid).Error; err != nil {
			return nil, errors.New("父类Id不存在！！！")
		}
	}
	if err := r.db.Where("cat_name = ?", name).First(&models.Category{}).Error; err == nil {
		return nil, errors.New("该分类已存在！！！")
	}
	var cate models.Category
	cate.CatPid = pid
	cate.CatLevel = level
	cate.CatName = name
	r.db.Debug().Create(&cate)
	result = map[string]interface{}{
		"cat_id":    cate.CatId,
		"cat_name":  cate.CatName,
		"cat_pid":   cate.CatPid,
		"cat_level": cate.CatLevel,
	}
	return

}
