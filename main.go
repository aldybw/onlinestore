package main

import (
	"log"
	"net/http"
	"onlinestore/auth"
	"onlinestore/handler"
	"onlinestore/helper"
	"onlinestore/product"
	shoppingcart "onlinestore/shoppingCart"
	"onlinestore/user"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/onlinestore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	productRepository := product.NewRepository(db)
	shoppingCartRepository := shoppingcart.NewRepository(db)

	userService := user.NewService(userRepository)
	productService := product.NewService(productRepository)
	shoppingCartService := shoppingcart.NewService(shoppingCartRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	productHandler := handler.NewProductHandler(productService)
	shoppingCartHandler := handler.NewShoppingCartHandler(shoppingCartService)
	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/products", productHandler.GetProducts)

	api.GET("/shoppingCarts", authMiddleware(authService, userService), shoppingCartHandler.GetShoppingCarts)
	api.POST("/shoppingCarts", authMiddleware(authService, userService), shoppingCartHandler.CreateShoppingCart)
	api.DELETE("/shoppingCarts", authMiddleware(authService, userService), shoppingCartHandler.DeleteShoppingCarts)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context)  {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}

		c.Set("currentUser", user)
	}
}

