package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kiki-ki/cal-app-go/api/handler"
)

func New() *gin.Engine {
	r := gin.Default()
	scheduleHandler := handler.NewScheduleHandler()
	scheduleGroup := r.Group("/schedules")
	{
		scheduleGroup.GET("", scheduleHandler.Index)
		scheduleGroup.POST("", scheduleHandler.Create)
	}
	return r
}
