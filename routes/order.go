package routes

import (
	"github.com/fazriegi/fga_assignment2/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewOrderRouter(db *gorm.DB, app *gin.Engine) {
	controller := controller.NewOrderController(db)

	app.POST("/orders", controller.CreateOrder)
	app.GET("/orders", controller.GetAllOrders)
	app.PUT("/orders/:id", controller.UpdateOrder)
	app.DELETE("/orders/:id", controller.DeleteOrder)
}
