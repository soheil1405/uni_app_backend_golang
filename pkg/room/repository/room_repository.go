package repository

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Create(room *models.Room) error
	Update(room *models.Room) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Room, error)
	List(filters *models.FetchRoomRequest) ([]*models.Room, error)
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) Create(room *models.Room) error {
	return r.db.Create(room).Error
}

func (r *roomRepository) Update(room *models.Room) error {
	return r.db.Save(room).Error
}

func (r *roomRepository) Delete(id database.PID) error {
	return r.db.Delete(&models.Room{}, id).Error
}

func (r *roomRepository) GetByID(id database.PID) (*models.Room, error) {
	var room models.Room
	err := r.db.Preload("Building").
		Preload("ClassSchedules").
		Preload("ClassSchedules.CourseInstance").
		Preload("ClassSchedules.CourseInstance.Course").
		Preload("ClassSchedules.CourseInstance.Teacher").
		First(&room, id).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) List(filters *models.FetchRoomRequest) ([]*models.Room, error) {
	var rooms []*models.Room
	query := r.db.Model(&models.Room{})

	if filters.BuildingID != 0 {
		query = query.Where("building_id = ?", filters.BuildingID)
	}

	if filters.Name != "" {
		query = query.Where("name LIKE ?", "%"+filters.Name+"%")
	}

	if filters.Capacity != 0 {
		query = query.Where("capacity = ?", filters.Capacity)
	}

	err := query.Preload("Building").
		Preload("ClassSchedules").
		Preload("ClassSchedules.CourseInstance").
		Preload("ClassSchedules.CourseInstance.Course").
		Preload("ClassSchedules.CourseInstance.Teacher").
		Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
