package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/auth/repository"
	routeRepositories "uni_app/pkg/route/repository"

	"github.com/labstack/echo/v4"
)

type AuthUsecase interface {
	CreateAuth(auth *models.AuthRules) error
	AuthEnforce(ctx echo.Context, req models.AuthRules, useCache bool) bool
	GetAuthByID(ctx echo.Context, ID database.PID, useCache bool) (*models.AuthRules, error)
	UpdateAuth(auth *models.AuthRules) error
	DeleteAuth(ID database.PID) error
	GetAllAuths() ([]models.AuthRules, error)
}

type authUsecase struct {
	repo      repositories.AuthRepository
	routeRepo routeRepositories.RouteRepository
}

func NewAuthUsecase(repo repositories.AuthRepository, routeRepo routeRepositories.RouteRepository) AuthUsecase {
	return &authUsecase{repo, routeRepo}
}

func (u *authUsecase) CreateAuth(auth *models.AuthRules) error {
	return u.repo.Create(auth)
}

func (u *authUsecase) GetAuthByID(ctx echo.Context, ID database.PID, useCache bool) (*models.AuthRules, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *authUsecase) UpdateAuth(auth *models.AuthRules) error {
	return u.repo.Update(auth)
}

func (u *authUsecase) DeleteAuth(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *authUsecase) GetAllAuths() ([]models.AuthRules, error) {
	return u.repo.GetAll()
}

func (u *authUsecase) AuthEnforce(ctx echo.Context, req models.AuthRules, useCache bool) bool {
	var (
		hasAccess = true
		route     = &models.Route{
			Url:    req.V4,
			Method: req.V5,
		}
		auth *models.AuthRules
		err  error
	)

	if route, err = u.routeRepo.FetchRoute(ctx, route); err != nil || route == nil || !route.ID.IsValid() {
		return hasAccess
	}

	req.V0 = route.RouteGroupID.String()
	if auth, err = u.repo.Enforce(ctx, req, useCache); err != nil || auth == nil || !auth.ID.IsValid() {
		hasAccess = false
	}

	return hasAccess
}
