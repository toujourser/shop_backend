package models

// 商品表
type Goods struct {
	GoodsId        int64   `gorm:"primary_key;AUTO_INCREMENT" json:"goods_id"`
	GoodsName      string  `gorm:"size:255" json:"goods_name"`         // 商品名称
	GoodsPrice     float64 `gorm:"default:0.00" json:"goods_price"`    // 商品价格
	GoodsNumber    int     `gorm:"default:0" json:"goods_number"`      // 商品数量
	GoodsWeight    float64 `gorm:"default:0" json:"goods_weight"`      // 商品重量
	CategoryId     int     `json:"category_id"`                        // 类型id
	GoodsIntroduce string  `gorm:"type:text" json:"goods_introduce"`   // 商品详情介绍
	GoodsBigLogo   string  `gorm:"default:''" json:"goods_big_logo"`   // 图片logo大图
	GoodsSmallLogo string  `gorm:"default:''" json:"goods_small_logo"` // 图片logo小图
	IsDel          int     `gorm:"default:0" json:"is_del"`            // 0:正常  1:删除
	AddTime        int64   `gorm:"size:11" json:"add_time"`
	UpdTime        int64   `gorm:"size:11" json:"upd_time"`
	DeleteTime     int64   `gorm:"size:11;NULL" json:"delete_time"`       // 软删除标志字段
	CatOneId       int     `gorm:"size:5;default:0" json:"cat_one_id"`    // 一级分类id
	CatTwoId       int     `gorm:"size:5;default:0" json:"cat_two_id"`    // 二级分类id
	CatThreeID     int     `gorm:"size:5;default:0" json:"cat_three_id"`  // 三级分类id
	HotNumber      int     `gorm:"size:11;default:0" json:"hot_Number"`   // 热卖数量
	IsPromote      int     `gorm:"size:5;default:0" json:"is_promote"`    // 是否促销
	GoodsState     int     `gorm:"size:11; default:0" json:"goods_state"` // 商品状态 0: 未通过 1: 审核中 2: 已审核
	//Category       *Category `gorm:"rel:fk" json:"cat_id"`              // 类型id
}

// 商品-属性关联表
type GoodsAttr struct {
	Id          int64   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	GoodsId     int64   `json:"goods_id"`                               // 商品id
	AttributeId int     `json:"attribute_id"`                           // 属性id
	AttrValue   string  `gorm:"type:text;default:''" json:"attr_value"` // 商品对应属性的值
	AddPrice    float64 `gorm:"NULL" json:"add_price"`                  // 该属性需要额外增加的价钱
	//Goods       *Goods     `gorm:"rel:fk" json:"goods_id"`                 // 商品id
	//Attribute   *Attribute `gorm:"rel:fk" json:"attr_id"`                  // 属性id
}

// 商品-相册关联表
type GoodsPics struct {
	PicsId  int    `gorm:"primary_key;AUTO_INCREMENT" json:"pics_id"`
	GoodsId int64  `json:"goods_id"`                   // 商品id
	PicsBig string `gorm:"default:''" json:"pics_big"` // 相册大图800*800
	PicsMid string `gorm:"default:''" json:"pics_mid"` // 相册中图350*350
	PicsSma string `gorm:"default:''" json:"pics_sma"` // 相册小图50*50
	//Goods   *Goods `gorm:"rel:fk" json:"goods_id"`
}

type GoodsCats struct {
	CatId      int64  `gorm:"primary_key;AUTO_INCREMENT" json:"cat_id"` // 分类id
	ParentId   int    `gorm:"size:11" json:"parent_id"`                 // 父级id
	CatName    string `gorm:"size:50" json:"cat_name"`                  // 分类名称
	IsShow     int    `gorm:"default:1" json:"is_show"`                 // 是否显示
	CatSort    int    `gorm:"default:0" json:"cat_sort"`                // 分类排序
	DataFlag   int    `gorm:"size:4" json:"data_flag"`                  // 数据标记
	CreateTime int    `gorm:"size:11" json:"create_time"`
}
