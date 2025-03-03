package dal

import (
	"gomall/model"

	"gorm.io/gorm"
)

// ProductDAL 定义数据访问接口
type ProductDAL interface {
	Create(product *model.Product) error
	GetByID(id int) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id int) error
	List(page, limit int) ([]model.Product, error)
}

type productDAL struct {
	db *gorm.DB
}

func NewProductDAL(db *gorm.DB) ProductDAL {
	return &productDAL{db: db}
}

// 创建商品
func (d *productDAL) Create(product *model.Product) error {
	return d.db.Create(product).Error
}

// 根据ID查询商品
func (d *productDAL) GetByID(id int) (*model.Product, error) {
	var product model.Product
	err := d.db.First(&product, id).Error
	return &product, err
}

// 更新商品
func (d *productDAL) Update(product *model.Product) error {
	return d.db.Save(product).Error
}

// 删除商品
func (d *productDAL) Delete(id int) error {
	return d.db.Delete(&model.Product{}, id).Error
}

// 分页查询商品列表
func (d *productDAL) List(page, limit int) ([]model.Product, error) {
	var products []model.Product
	offset := (page - 1) * limit
	err := d.db.Offset(offset).Limit(limit).Find(&products).Error
	return products, err
}
