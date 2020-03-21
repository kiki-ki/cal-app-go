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

func (d *db) CreateSchedule(schedule model.Schedule) error {
	if e := d.orm.Create(&schedule).Error; e != nil {
		return e
	}
	return nil
}

func (d *db) UpdateScheduleByID(id int, schedule model.Schedule) error {
	var schedule model.Schedule
	if e != d.orm.Model(&schedule).Updates(Schedule{Name: schedule.Name}); e != nil {

	}
}

func (d *db) DeleteScheduleByID(id int) error {
	var schedule model.Schedule
	if e := d.orm.Where("id = ?", id).Delete(&schedule).Error; e != nil {
		return e
	}
	return nil
}
