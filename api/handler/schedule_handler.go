package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kiki-ki/cal-app-go/src/database"
	"github.com/kiki-ki/cal-app-go/src/model"
)

type ScheduleHandler interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	// New(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
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

func (handler scheduleHandler) Show(c *gin.Context) {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}
	schedule, e := database.DB.FindScheduleByID(id)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}
	c.JSON(http.StatusOK, schedule)
	return
}

func (handler scheduleHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "templates/schedules/new.html", gin.H{})
}

func (handler scheduleHandler) Create(c *gin.Context) {
	var schedule model.Schedule
	if e := c.BindJSON(&schedule); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}
	if e := database.DB.CreateSchedule(schedule); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}
	c.JSON(http.StatusCreated, schedule)
}

func (handler scheduleHandler) Update(c *gin.Context) {
	var schedule model.Schedule

	if id, e := strconv.Atoi(c.Param("id")); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}
	if e := c.BindJSON(&schedule); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}
	if e := database.DB.UpdateScheduleByID(id, schedule); e != nil {

	}
}

func (handler scheduleHandler) Destroy(c *gin.Context) {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}
	if e := database.DB.DeleteScheduleByID(id); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
