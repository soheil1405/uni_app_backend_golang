package repository

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type BuildingRepository interface {
	Create(building *models.Building) error
	Update(building *models.Building) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Building, error)
	List(filters *models.FetchBuildingRequest) ([]*models.Building, error)
}

type buildingRepository struct {
	db *gorm.DB
}

func NewBuildingRepository(db *gorm.DB) BuildingRepository {
	return &buildingRepository{db: db}
}

func (r *buildingRepository) Create(building *models.Building) error {
	return r.db.Create(building).Error
}

func (r *buildingRepository) Update(building *models.Building) error {
	return r.db.Save(building).Error
}

func (r *buildingRepository) Delete(id database.PID) error {
	return r.db.Delete(&models.Building{}, id).Error
}

func (r *buildingRepository) GetByID(id database.PID) (*models.Building, error) {
	var building models.Building
	err := r.db.Preload("Rooms").
		Preload("Rooms.ClassSchedules").
		Preload("Rooms.ClassSchedules.CourseInstance").
		Preload("Rooms.ClassSchedules.CourseInstance.Course").
		Preload("Rooms.ClassSchedules.CourseInstance.Teacher").
		First(&building, id).Error
	if err != nil {
		return nil, err
	}
	return &building, nil
}

func (r *buildingRepository) List(filters *models.FetchBuildingRequest) ([]*models.Building, error) {
	var buildings []*models.Building
	query := r.db.Model(&models.Building{})

	if filters.Name != "" {
		query = query.Where("name LIKE ?", "%"+filters.Name+"%")
	}

	if filters.Address != "" {
		query = query.Where("address LIKE ?", "%"+filters.Address+"%")
	}

	err := query.Preload("Rooms").
		Preload("Rooms.ClassSchedules").
		Preload("Rooms.ClassSchedules.CourseInstance").
		Preload("Rooms.ClassSchedules.CourseInstance.Course").
		Preload("Rooms.ClassSchedules.CourseInstance.Teacher").
		Find(&buildings).Error
	if err != nil {
		return nil, err
	}
	return buildings, nil
}
