package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kiki-ki/cal-app-go/src/database"
	"github.com/kiki-ki/cal-app-go/src/model"
)

type ScheduleHandler interface {
	Index(c *gin.Context)
	New(c *gin.Context)
	Create(c *gin.Context)
}

type scheduleHandler struct{}

func NewScheduleHandler() ScheduleHandler {
	var handler ScheduleHandler
	handler = scheduleHandler{}
	return handler
}

func (handler scheduleHandler) Index(c *gin.Context) {
	schedules, _ := database.DB.FindAllSchedules()
	c.JSON(http.StatusOK, schedules)
}

func (handler scheduleHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "schedules/new.html", gin.H{})
}

func (handler scheduleHandler) Create(c *gin.Context) {
	timeNow := time.Now()
	schedule := model.Schedule{Title: "aaa", FromDate: &timeNow, ToDate: &timeNow}
	database.DB.CreateSchedule(schedule)
	c.JSON(http.StatusCreated, schedule)
}
