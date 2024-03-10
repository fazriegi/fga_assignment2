package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fazriegi/fga_assignment2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderController struct {
	db *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{db}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order

	if err := ctx.BindJSON(&order); err != nil {
		log.Printf("error: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Println("order", order)
	if err := c.db.Create(&order).Error; err != nil {
		log.Printf("error creating order: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create order"})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	var orders []models.Order

	if err := c.db.Preload("Items").Find(&orders).Error; err != nil {
		log.Printf("error fetching orders: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch orders"})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	var order models.Order
	orderId := ctx.Param("id")

	if err := c.db.Preload("Items").First(&order, orderId).Error; err != nil {
		log.Printf("error fetching order: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		return
	}

	if err := ctx.BindJSON(&order); err != nil {
		log.Printf("error: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := c.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error; err != nil {
		log.Printf("error update order: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete order"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	var order models.Order
	orderId := ctx.Param("id")

	if err := c.db.First(&order, orderId).Error; err != nil {
		log.Printf("error fetching order: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := c.db.Select(clause.Associations).Delete(&order).Error; err != nil {
		log.Printf("error deleting order: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete order"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
