package usecase

import (
	"errors"
	"time"
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/request/repository"

	"github.com/labstack/echo/v4"
)

type RequestUsecase interface {
	CreateRequest(request *models.FetchRequest) error
	GetRequestByID(ctx echo.Context, ID database.PID, useCache bool) (*models.FetchRequest, error)
	UpdateRequest(request *models.FetchRequest) error
	DeleteRequest(ID database.PID) error
	GetAllRequests() ([]models.FetchRequest, error)
	GetRequestsByUserID(userID database.PID) ([]models.FetchRequest, error)
	GetRequestsByStatus(status string) ([]models.FetchRequest, error)
	GetRequestsByDateRange(startDate, endDate time.Time) ([]models.FetchRequest, error)
}

type requestUsecase struct {
	repo repository.RequestRepository
}

func NewRequestUsecase(repo repository.RequestRepository) RequestUsecase {
	return &requestUsecase{repo}
}

func (u *requestUsecase) CreateRequest(request *models.FetchRequest) error {
	// Validate required fields
	if request.UserID == 0 {
		return errors.New("user ID is required")
	}
	if request.Status == "" {
		request.Status = "pending" // Set default status
	}
	if request.CreatedAt.IsZero() {
		request.CreatedAt = time.Now()
	}

	// Validate status
	validStatuses := map[string]bool{
		"pending":   true,
		"approved":  true,
		"rejected":  true,
		"completed": true,
	}
	if !validStatuses[request.Status] {
		return errors.New("invalid status")
	}

	return u.repo.Create(request)
}

func (u *requestUsecase) GetRequestByID(ctx echo.Context, ID database.PID, useCache bool) (*models.FetchRequest, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *requestUsecase) UpdateRequest(request *models.FetchRequest) error {
	// Validate required fields
	if request.UserID == 0 {
		return errors.New("user ID is required")
	}

	// Validate status
	validStatuses := map[string]bool{
		"pending":   true,
		"approved":  true,
		"rejected":  true,
		"completed": true,
	}
	if !validStatuses[request.Status] {
		return errors.New("invalid status")
	}

	// Update timestamp
	request.UpdatedAt = time.Now()

	return u.repo.Update(request)
}

func (u *requestUsecase) DeleteRequest(ID database.PID) error {
	// Check if request exists and can be deleted
	request, err := u.repo.GetByID(nil, ID, false)
	if err != nil {
		return err
	}

	// Only allow deletion of pending requests
	if request.Status != "pending" {
		return errors.New("can only delete pending requests")
	}

	return u.repo.Delete(ID)
}

func (u *requestUsecase) GetAllRequests() ([]models.FetchRequest, error) {
	return u.repo.GetAll()
}

func (u *requestUsecase) GetRequestsByUserID(userID database.PID) ([]models.FetchRequest, error) {
	allRequests, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredRequests []models.FetchRequest
	for _, request := range allRequests {
		if request.UserID == userID {
			filteredRequests = append(filteredRequests, request)
		}
	}

	return filteredRequests, nil
}

func (u *requestUsecase) GetRequestsByStatus(status string) ([]models.FetchRequest, error) {
	allRequests, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredRequests []models.FetchRequest
	for _, request := range allRequests {
		if request.Status == status {
			filteredRequests = append(filteredRequests, request)
		}
	}

	return filteredRequests, nil
}

func (u *requestUsecase) GetRequestsByDateRange(startDate, endDate time.Time) ([]models.FetchRequest, error) {
	allRequests, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredRequests []models.FetchRequest
	for _, request := range allRequests {
		if request.CreatedAt.After(startDate) && request.CreatedAt.Before(endDate) {
			filteredRequests = append(filteredRequests, request)
		}
	}

	return filteredRequests, nil
} 