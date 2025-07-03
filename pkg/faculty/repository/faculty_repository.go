package repository

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type FacultyRepository interface {
	Create(faculty *models.Faculty) error
	Update(faculty *models.Faculty) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Faculty, error)
	List(filters *models.FetchFacultyRequest) ([]*models.Faculty, error)
	AddDepartment(facultyID, departmentID database.PID) error
	RemoveDepartment(facultyID, departmentID database.PID) error
	AddTeacher(facultyID, teacherID database.PID) error
	RemoveTeacher(facultyID, teacherID database.PID) error
	AddStaff(facultyID, userID database.PID) error
	RemoveStaff(facultyID, userID database.PID) error
}

type facultyRepository struct {
	db *gorm.DB
}

func NewFacultyRepository(db *gorm.DB) FacultyRepository {
	return &facultyRepository{db: db}
}

func (r *facultyRepository) Create(faculty *models.Faculty) error {
	return r.db.Create(faculty).Error
}

func (r *facultyRepository) Update(faculty *models.Faculty) error {
	return r.db.Save(faculty).Error
}

func (r *facultyRepository) Delete(id database.PID) error {
	return r.db.Delete(&models.Faculty{}, id).Error
}

func (r *facultyRepository) GetByID(id database.PID) (*models.Faculty, error) {
	var faculty models.Faculty
	err := r.db.Preload("University").
		Preload("Floors").
		Preload("Departments").
		Preload("Students").
		Preload("Teachers").
		Preload("Staff").
		First(&faculty, id).Error
	if err != nil {
		return nil, err
	}
	return &faculty, nil
}

func (r *facultyRepository) List(filters *models.FetchFacultyRequest) ([]*models.Faculty, error) {
	var faculties []*models.Faculty
	query := r.db.Model(&models.Faculty{})

	if filters.UniversityID.IsValid() {
		query = query.Where("university_id = ?", filters.UniversityID)
	}

	if filters.Name != "" {
		query = query.Where("name LIKE ?", "%"+filters.Name+"%")
	}

	err := query.Preload("University").
		Preload("Floors").
		Preload("Departments").
		Preload("Students").
		Preload("Teachers").
		Preload("Staff").
		Find(&faculties).Error
	if err != nil {
		return nil, err
	}
	return faculties, nil
}

func (r *facultyRepository) AddDepartment(facultyID, departmentID database.PID) error {
	return r.db.Model(&models.Faculty{}).Where("id = ?", facultyID).
		Association("Departments").Append(&models.Department{Model: database.Model{ID: departmentID}})
}

func (r *facultyRepository) RemoveDepartment(facultyID, departmentID database.PID) error {
	return r.db.Model(&models.Faculty{}).Where("id = ?", facultyID).
		Association("Departments").Delete(&models.Department{Model: database.Model{ID: departmentID}})
}

func (r *facultyRepository) AddTeacher(facultyID, teacherID database.PID) error {
	return r.db.Model(&models.Faculty{}).Where("id = ?", facultyID).
		Association("Teachers").Append(&models.Teacher{Model: database.Model{ID: teacherID}})
}

func (r *facultyRepository) RemoveTeacher(facultyID, teacherID database.PID) error {
	return r.db.Model(&models.Faculty{}).Where("id = ?", facultyID).
		Association("Teachers").Delete(&models.Teacher{Model: database.Model{ID: teacherID}})
}

func (r *facultyRepository) AddStaff(facultyID, userID database.PID) error {
	return r.db.Model(&models.Faculty{}).Where("id = ?", facultyID).
		Association("Staff").Append(&models.User{Model: database.Model{ID: userID}})
}

func (r *facultyRepository) RemoveStaff(facultyID, userID database.PID) error {
	return r.db.Model(&models.Faculty{}).Where("id = ?", facultyID).
		Association("Staff").Delete(&models.User{Model: database.Model{ID: userID}})
}
