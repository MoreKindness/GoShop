package product

import (
	"gomall/model"
	"gomall/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// GetProductByID 创建商品（HTTP: POST /products）
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "无效的请求数据",
			Data:    nil,
		})
		return
	}

	if err := h.productService.SaveProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "服务端处理失败",
			Data:    nil,
		})
		return
	}

	// 返回标准化成功响应
	c.JSON(http.StatusCreated, ApiResponse{
		Code:    http.StatusOK,
		Message: "商品创建成功",
	})
}

// Searchproducs 查询单个商品（HTTP: GET /products/:id）
func (h *ProductHandler) SearchProducts(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	log.Printf("id: %d", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的商品ID"})
		return
	}

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "商品不存在"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct 更新商品（HTTP: PUT /products/:id）
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的商品ID"})
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	product.ID = id // 确保 ID 一致性

	if err := h.productService.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct 删除商品（HTTP: DELETE /products/:id）
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的商品ID"})
		return
	}

	if err := h.productService.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// ListProducts 分页查询商品列表（HTTP: GET /products）
func (h *ProductHandler) ListProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, err := h.productService.ListProducts(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
