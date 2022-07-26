package handler

import (
	"net/http"
	"onlinestore/helper"
	shoppingcart "onlinestore/shoppingCart"
	"onlinestore/user"

	"github.com/gin-gonic/gin"
)

type shoppingCartHandler struct {
	service shoppingcart.Service
}

func NewShoppingCartHandler(service shoppingcart.Service) *shoppingCartHandler {
	return &shoppingCartHandler{service}
}

func (h *shoppingCartHandler) GetShoppingCarts(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	userID := currentUser.ID

	shoppingCarts, err := h.service.GetShoppingCarts(userID)
	if err != nil {
		response := helper.APIResponse("Error to get shopping carts", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	response := helper.APIResponse("List of shopping carts", http.StatusOK, "success", shoppingcart.FormatShoppingCarts(shoppingCarts, userID))
	c.JSON(http.StatusOK, response)
}

func (h *shoppingCartHandler) CreateShoppingCart(c *gin.Context) {
	var input shoppingcart.CreateShoppingCartInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create shopping cart", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	userID := currentUser.ID

	newShoppingCart, err := h.service.CreateShoppingCart(input, userID)
	if err != nil {
		response := helper.APIResponse("Failed to create shopping cart", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create shopping cart", http.StatusOK, "success", shoppingcart.FormatShoppingCart(newShoppingCart, userID))
	c.JSON(http.StatusOK, response)
}

func (h *shoppingCartHandler) DeleteShoppingCarts(c *gin.Context) {
	var input shoppingcart.GetShoppingCartDetailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete shopping cart", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// currentUser := c.MustGet("currentUser").(user.User)

	// userID := currentUser.ID

	h.service.DeleteShoppingCart(input)
	if err != nil {
		response := helper.APIResponse("Error to delete shopping carts", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	response := helper.APIResponse("Shopping cart deleted", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}