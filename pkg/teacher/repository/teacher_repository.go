package repository

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	Create(teacher *models.Teacher) error
	Update(teacher *models.Teacher) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Teacher, error)
	List(filters *models.FetchTeacherRequest) ([]*models.Teacher, error)
	AssignOffice(teacherID, roomID database.PID) error
	AssignCourse(teacherID, courseInstanceID database.PID) error
	RemoveCourse(teacherID, courseInstanceID database.PID) error
}

type teacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{db: db}
}

func (r *teacherRepository) Create(teacher *models.Teacher) error {
	return r.db.Create(teacher).Error
}

func (r *teacherRepository) Update(teacher *models.Teacher) error {
	return r.db.Save(teacher).Error
}

func (r *teacherRepository) Delete(id database.PID) error {
	return r.db.Delete(&models.Teacher{}, id).Error
}

func (r *teacherRepository) GetByID(id database.PID) (*models.Teacher, error) {
	var teacher models.Teacher
	err := r.db.Preload("CourseInstances").
		Preload("CourseInstances.Course").
		Preload("CourseInstances.ClassSchedules").
		Preload("CourseInstances.ClassSchedules.Room").
		Preload("CourseInstances.ClassSchedules.Room.Building").
		First(&teacher, id).Error
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *teacherRepository) List(filters *models.FetchTeacherRequest) ([]*models.Teacher, error) {
	var teachers []*models.Teacher
	query := r.db.Model(&models.Teacher{})

	if filters.Name != "" {
		query = query.Where("name LIKE ?", "%"+filters.Name+"%")
	}

	if filters.Email != "" {
		query = query.Where("email LIKE ?", "%"+filters.Email+"%")
	}

	if filters.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+filters.Phone+"%")
	}

	err := query.Preload("CourseInstances").
		Preload("CourseInstances.Course").
		Preload("CourseInstances.ClassSchedules").
		Preload("CourseInstances.ClassSchedules.Room").
		Preload("CourseInstances.ClassSchedules.Room.Building").
		Find(&teachers).Error
	if err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *teacherRepository) AssignOffice(teacherID, roomID database.PID) error {
	return r.db.Model(&models.Teacher{}).Where("id = ?", teacherID).
		Update("office_room_id", roomID).Error
}

func (r *teacherRepository) AssignCourse(teacherID, courseInstanceID database.PID) error {
	return r.db.Model(&models.Teacher{}).Where("id = ?", teacherID).
		Association("Courses").Append(&models.CourseInstance{Model: database.Model{ID: courseInstanceID}})
}

func (r *teacherRepository) RemoveCourse(teacherID, courseInstanceID database.PID) error {
	return r.db.Model(&models.Teacher{}).Where("id = ?", teacherID).
		Association("Courses").Delete(&models.CourseInstance{Model: database.Model{ID: courseInstanceID}})
}
