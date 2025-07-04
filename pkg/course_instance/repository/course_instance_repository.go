package repository

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type CourseInstanceRepository interface {
	Create(courseInstance *models.CourseInstance) error
	Update(courseInstance *models.CourseInstance) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.CourseInstance, error)
	List(filters *models.FetchCourseInstanceRequest) ([]*models.CourseInstance, error)
	AddStudent(courseInstanceID, studentID database.PID) error
	RemoveStudent(courseInstanceID, studentID database.PID) error
}

type courseInstanceRepository struct {
	db *gorm.DB
}

func NewCourseInstanceRepository(db *gorm.DB) CourseInstanceRepository {
	return &courseInstanceRepository{db: db}
}

func (r *courseInstanceRepository) Create(courseInstance *models.CourseInstance) error {
	return r.db.Create(courseInstance).Error
}

func (r *courseInstanceRepository) Update(courseInstance *models.CourseInstance) error {
	return r.db.Save(courseInstance).Error
}

func (r *courseInstanceRepository) Delete(id database.PID) error {
	return r.db.Delete(&models.CourseInstance{}, id).Error
}

func (r *courseInstanceRepository) GetByID(id database.PID) (*models.CourseInstance, error) {
	var courseInstance models.CourseInstance
	err := r.db.Preload("Course").
		Preload("Teacher").
		Preload("Students").
		Preload("Room").
		Preload("Faculty").
		First(&courseInstance, id).Error
	if err != nil {
		return nil, err
	}
	return &courseInstance, nil
}

func (r *courseInstanceRepository) List(filters *models.FetchCourseInstanceRequest) ([]*models.CourseInstance, error) {
	var courseInstances []*models.CourseInstance
	query := r.db.Model(&models.CourseInstance{})

	if filters.CourseID != 0 {
		query = query.Where("course_id = ?", filters.CourseID)
	}

	if filters.TeacherID != 0 {
		query = query.Where("teacher_id = ?", filters.TeacherID)
	}

	err := query.Preload("Course").
		Preload("Teacher").
		Preload("Students").
		Preload("Room").
		Preload("Faculty").
		Find(&courseInstances).Error
	if err != nil {
		return nil, err
	}
	return courseInstances, nil
}

func (r *courseInstanceRepository) AddStudent(courseInstanceID, studentID database.PID) error {
	return r.db.Model(&models.CourseInstance{Model: database.Model{ID: courseInstanceID}}).
		Association("Students").Append(&models.Student{Model: database.Model{ID: studentID}})
}

func (r *courseInstanceRepository) RemoveStudent(courseInstanceID, studentID database.PID) error {
	return r.db.Model(&models.CourseInstance{Model: database.Model{ID: courseInstanceID}}).
		Association("Students").Delete(&models.Student{Model: database.Model{ID: studentID}})
}
