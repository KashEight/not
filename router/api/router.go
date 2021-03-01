package api

import (
	"github.com/KashEight/not/middleware"
	"github.com/KashEight/not/repos"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	repo repos.Repo
}

func Init(engine *gin.Engine, db *gorm.DB) {
	h := newHandler(db)
	apiGroup := engine.Group("/api")
	{
		noteGroup := apiGroup.Group("/note")
		{
			noteGroup.GET("/", h.getAllNote())
			noteGroup.POST("/", h.createNote())
			noteUUIDGroup := noteGroup.Group("/:uuid", middleware.SetUUIDKey())
			{
				noteUUIDGroup.GET("/", h.getNote())
				noteUUIDGroup.PATCH("/", h.updateNote())
				noteUUIDGroup.DELETE("/", h.deleteNote())
			}
		}
	}
}

func newHandler(db *gorm.DB) *handler {
	repo := repos.NewRepository(db)
	h := &handler{
		repo: repo,
	}
	return h
}
