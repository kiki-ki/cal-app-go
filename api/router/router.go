package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kiki-ki/cal-app-go/api/handler"
)

func New() *gin.Engine {
	router := gin.Default()
	scheduleHandler := handler.NewScheduleHandler()
	scheduleGroup := router.Group("/schedules")
	{
		scheduleGroup.GET("", scheduleHandler.Index)
		scheduleGroup.GET("/:id", scheduleHandler.Show)
		//scheduleGroup.GET("/new", scheduleHandler.New)
		scheduleGroup.POST("", scheduleHandler.Create)
		scheduleGroup.DELETE("/:id", scheduleHandler.Destroy)
	}
	return router
}
