package service

import (
	"errors"
	"gomall/dal"
	"gomall/dal/mysql"
	"gomall/model"
)

// ProductService 定义业务逻辑接口
type ProductService interface {
	GetProductByID(id int) (*model.Product, error)
	SearchProducts(id int) (*model.Product, error)
	SaveProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(id int) error
	ListProducts(page, limit int) ([]model.Product, error)
	ListProductsByCategory(category string, page, limit int) ([]model.Product, error)
}

type productService struct {
	productDAL dal.ProductDAL
}

func NewProductService() ProductService {
	dal := dal.NewProductDAL(mysql.DB)
	return &productService{
		productDAL: dal,
	}
}

func (s *productService) GetProductByID(id int) (*model.Product, error) {
	// 调用 DAL 层方法
	return s.productDAL.GetByID(id)
}

// 创建商品（业务逻辑：名称不能为空，价格必须大于0）
func (s *productService) SaveProduct(product *model.Product) error {
	if product.Name == "" {
		return errors.New("商品名称不能为空")
	}
	if product.Price <= 0 {
		return errors.New("商品价格必须大于0")
	}
	return s.productDAL.Create(product)
}

// 查询单个商品
func (s *productService) SearchProducts(id int) (*model.Product, error) {
	return s.productDAL.GetByID(id)
}

// 更新商品
func (s *productService) UpdateProduct(product *model.Product) error {
	return s.productDAL.Update(product)
}

// 删除商品
func (s *productService) DeleteProduct(id int) error {
	return s.productDAL.Delete(id)
}

// 分页查询商品列表
func (s *productService) ListProducts(page, limit int) ([]model.Product, error) {
	return s.productDAL.List(page, limit)
}

func (s *productService) ListProductsByCategory(category string, page, limit int) ([]model.Product, error) {
	return s.productDAL.ListProductsByCategory(category, page, limit)
}
