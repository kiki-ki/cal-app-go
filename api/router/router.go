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
		scheduleGroup.GET("/new", scheduleHandler.New)
		scheduleGroup.POST("", scheduleHandler.Create)
	}
	return router
}
