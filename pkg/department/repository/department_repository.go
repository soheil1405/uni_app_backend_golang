package repository

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(department *models.Department) error
	Update(department *models.Department) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Department, error)
	List(filters *models.FetchDepartmentRequest) ([]*models.Department, error)
	SetHead(departmentID, teacherID database.PID) error
	AddTeacher(departmentID, teacherID database.PID) error
	RemoveTeacher(departmentID, teacherID database.PID) error
	AddMajor(departmentID, majorID database.PID) error
	RemoveMajor(departmentID, majorID database.PID) error
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) Create(department *models.Department) error {
	return r.db.Create(department).Error
}

func (r *departmentRepository) Update(department *models.Department) error {
	return r.db.Save(department).Error
}

func (r *departmentRepository) Delete(id database.PID) error {
	return r.db.Delete(&models.Department{}, id).Error
}

func (r *departmentRepository) GetByID(id database.PID) (*models.Department, error) {
	var department models.Department
	err := r.db.Preload("Faculty").
		Preload("Head").
		Preload("Teachers").
		Preload("Majors").
		First(&department, id).Error
	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *departmentRepository) List(filters *models.FetchDepartmentRequest) ([]*models.Department, error) {
	var departments []*models.Department
	query := r.db.Model(&models.Department{})

	if filters.FacultyID.IsValid() {
		query = query.Where("faculty_id = ?", filters.FacultyID)
	}

	if filters.Name != "" {
		query = query.Where("name LIKE ?", "%"+filters.Name+"%")
	}

	err := query.Preload("Faculty").
		Preload("Head").
		Preload("Teachers").
		Preload("Majors").
		Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *departmentRepository) SetHead(departmentID, teacherID database.PID) error {
	return r.db.Model(&models.Department{}).Where("id = ?", departmentID).
		Update("head_id", teacherID).Error
}

func (r *departmentRepository) AddTeacher(departmentID, teacherID database.PID) error {
	return r.db.Model(&models.Department{}).Where("id = ?", departmentID).
		Association("Teachers").Append(&models.Teacher{Model: database.Model{ID: teacherID}})
}

func (r *departmentRepository) RemoveTeacher(departmentID, teacherID database.PID) error {
	return r.db.Model(&models.Department{}).Where("id = ?", departmentID).
		Association("Teachers").Delete(&models.Teacher{Model: database.Model{ID: teacherID}})
}

func (r *departmentRepository) AddMajor(departmentID, majorID database.PID) error {
	return r.db.Model(&models.Department{}).Where("id = ?", departmentID).
		Association("Majors").Append(&models.Major{Model: database.Model{ID: majorID}})
}

func (r *departmentRepository) RemoveMajor(departmentID, majorID database.PID) error {
	return r.db.Model(&models.Department{}).Where("id = ?", departmentID).
		Association("Majors").Delete(&models.Major{Model: database.Model{ID: majorID}})
}
