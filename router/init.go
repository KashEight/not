package router

import (
	"github.com/KashEight/not/router/api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(engine *gin.Engine, db *gorm.DB) {
	api.Init(engine, db)
}
