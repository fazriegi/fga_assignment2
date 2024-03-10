package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, app *gin.Engine) {
	NewOrderRouter(db, app)
}
