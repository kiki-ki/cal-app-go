package database

import (
	"github.com/kiki-ki/cal-app-go/src/model"
)

// errorを返すかは要検討
func (d *db) FindAllSchedules() ([]model.Schedule, error) {
	var schedules []model.Schedule
	d.orm.Find(&schedules)

	return schedules, nil
}

func (d *db) FindScheduleByID(id int) (model.Schedule, error) {
	var schedule model.Schedule
	if e := d.orm.Find(&schedule, id).Error; e != nil {
		return schedule, e
	}
	return schedule, nil
}

func (d *db) CreateSchedule(schedule model.Schedule) {
	d.orm.Create(&schedule)
}
