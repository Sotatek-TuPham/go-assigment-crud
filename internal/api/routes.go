// internal/api/routes.go
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hauchu1196/ecombase/internal/api/handlers"
	"github.com/hauchu1196/ecombase/internal/repository"
	"github.com/hauchu1196/ecombase/internal/service"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Order routes
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("", orderHandler.CreateOrder)
		orderRoutes.GET("/:id", orderHandler.GetOrder)
		orderRoutes.GET("", orderHandler.ListOrders)
		orderRoutes.PUT("/:id", orderHandler.UpdateOrder)
		orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
	}

	bookRoutes := router.Group("/books")
	{
		bookRoutes.POST("", bookHandler.CreateBook)
		bookRoutes.GET("/:id", bookHandler.GetBook)
		bookRoutes.PUT("/:id", bookHandler.UpdateBook)
		bookRoutes.DELETE("/:id", bookHandler.DeleteBook)
	}
}
