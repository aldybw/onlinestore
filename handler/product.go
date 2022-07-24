package handler

import (
	"net/http"
	"onlinestore/helper"
	"onlinestore/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Query("category_id"))

	products, err := h.service.GetProducts(categoryID)
	if err != nil {
		response := helper.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of products", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}