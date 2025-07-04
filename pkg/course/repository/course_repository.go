package repository

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *models.Course) error
	Update(course *models.Course) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Course, error)
	List(filters *models.FetchCourseRequest) ([]*models.Course, error)
	AddPrerequisite(courseID, prerequisiteID database.PID) error
	RemovePrerequisite(courseID, prerequisiteID database.PID) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(course *models.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) Update(course *models.Course) error {
	return r.db.Save(course).Error
}

func (r *courseRepository) Delete(id database.PID) error {
	return r.db.Delete(&models.Course{}, id).Error
}

func (r *courseRepository) GetByID(id database.PID) (*models.Course, error) {
	var course models.Course
	err := r.db.Preload("CourseInstances").
		Preload("CourseInstances.Teacher").
		Preload("CourseInstances.ClassSchedules").
		Preload("CourseInstances.ClassSchedules.Room").
		Preload("CourseInstances.ClassSchedules.Room.Building").
		First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) List(filters *models.FetchCourseRequest) ([]*models.Course, error) {
	var courses []*models.Course
	query := r.db.Model(&models.Course{})

	if filters.Name != "" {
		query = query.Where("name LIKE ?", "%"+filters.Name+"%")
	}

	err := query.Preload("CourseInstances").
		Preload("CourseInstances.Teacher").
		Preload("CourseInstances.ClassSchedules").
		Preload("CourseInstances.ClassSchedules.Room").
		Preload("CourseInstances.ClassSchedules.Room.Building").
		Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *courseRepository) AddPrerequisite(courseID, prerequisiteID database.PID) error {
	prerequisite := &models.CoursePrerequisite{
		CourseID:       courseID,
		PrerequisiteID: prerequisiteID,
	}
	return r.db.Create(prerequisite).Error
}

func (r *courseRepository) RemovePrerequisite(courseID, prerequisiteID database.PID) error {
	return r.db.Where("course_id = ? AND prerequisite_id = ?", courseID, prerequisiteID).
		Delete(&models.CoursePrerequisite{}).Error
}
