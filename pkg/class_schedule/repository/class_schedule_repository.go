package repository

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type ClassScheduleRepository interface {
	Create(classSchedule *models.ClassSchedule) error
	Update(classSchedule *models.ClassSchedule) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.ClassSchedule, error)
	List(filters *models.FetchClassScheduleRequest) ([]*models.ClassSchedule, error)
}

type classScheduleRepository struct {
	db *gorm.DB
}

func NewClassScheduleRepository(db *gorm.DB) ClassScheduleRepository {
	return &classScheduleRepository{db: db}
}

func (r *classScheduleRepository) Create(classSchedule *models.ClassSchedule) error {
	return r.db.Create(classSchedule).Error
}

func (r *classScheduleRepository) Update(classSchedule *models.ClassSchedule) error {
	return r.db.Save(classSchedule).Error
}

func (r *classScheduleRepository) Delete(id database.PID) error {
	return r.db.Delete(&models.ClassSchedule{}, id).Error
}

func (r *classScheduleRepository) GetByID(id database.PID) (*models.ClassSchedule, error) {
	var classSchedule models.ClassSchedule
	err := r.db.Preload("CourseInstance").
		Preload("CourseInstance.Course").
		Preload("CourseInstance.Teacher").
		Preload("Room").
		Preload("Room.Building").
		First(&classSchedule, id).Error
	if err != nil {
		return nil, err
	}
	return &classSchedule, nil
}

func (r *classScheduleRepository) List(filters *models.FetchClassScheduleRequest) ([]*models.ClassSchedule, error) {
	var classSchedules []*models.ClassSchedule
	query := r.db.Model(&models.ClassSchedule{})

	if filters.CourseInstanceID != 0 {
		query = query.Where("course_instance_id = ?", filters.CourseInstanceID)
	}

	if filters.RoomID != 0 {
		query = query.Where("room_id = ?", filters.RoomID)
	}

	if filters.DayOfWeek != "" {
		query = query.Where("day_of_week = ?", filters.DayOfWeek)
	}

	if filters.StartTime != "" {
		query = query.Where("start_time = ?", filters.StartTime)
	}

	if filters.EndTime != "" {
		query = query.Where("end_time = ?", filters.EndTime)
	}

	err := query.Preload("CourseInstance").
		Preload("CourseInstance.Course").
		Preload("CourseInstance.Teacher").
		Preload("Room").
		Preload("Room.Building").
		Find(&classSchedules).Error
	if err != nil {
		return nil, err
	}
	return classSchedules, nil
}
