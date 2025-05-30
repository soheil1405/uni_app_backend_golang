package middleware

import (
	"net/http"
	"regexp"
	"strings"
	"time"
	"uni_app/database"
	"uni_app/models"
	_authRepository "uni_app/pkg/auth/repository"
	_authUc "uni_app/pkg/auth/usecase"
	_routeRepository "uni_app/pkg/route/repository"
	_studentRepository "uni_app/pkg/student/repository"
	_userRepository "uni_app/pkg/user/repository"
	"uni_app/services/env"
	"uni_app/utils/helpers"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

// GoMiddleware ...
type GoMiddleware struct {
	userRepo          _userRepository.UserRepository
	studentRepository _studentRepository.StudentRepository
	authUcecase       _authUc.AuthUsecase
	config            *env.Config
}

// InitMiddleware ...
func InitMiddleware(e *echo.Group, db *gorm.DB, config *env.Config) *GoMiddleware {
	authRepo := _authRepository.NewAuthRepository(db)
	routeRepo := _routeRepository.NewRouteRepository(db)
	authUC := _authUc.NewAuthUsecase(authRepo, routeRepo)
	return &GoMiddleware{
		userRepo:          _userRepository.NewUserRepository(db),
		studentRepository: _studentRepository.NewStudentRepository(db),
		authUcecase:       authUC,
		config:            config,
	}
}

var (
	callbackRegexp = regexp.MustCompile("^/api/v1(/.+)?/payments/.*/callback/.*")
)

// SkipSetContext skips setting context for specified routes
func (m *GoMiddleware) SkipSetContext(skipper func(echo.Context) bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if skipper(c) {
				return next(c)
			}
			// Continue with setting context if not skipped
			return m.SetContext(next)(c)
		}
	}
}

func RegisterSkipper(ctx echo.Context) bool {
	if strings.Contains(ctx.Path(), "/api/v1/users/register") ||
		strings.Contains(ctx.Path(), "/api/v1/users/login") ||
		strings.Contains(ctx.Path(), "/api/v1/unis") ||
		strings.Contains(ctx.Path(), "/api/v1/majors") ||
		callbackRegexp.MatchString(ctx.Path()) {
		return true
	}
	return false
}

func cachSkipper(ctx echo.Context) bool {
	path := ctx.Path()

	return (strings.Index(path, "/api/v1/auth") == 0)
}

func validateToken(token *jwt.Token, signingKey string) error {
	if !token.Valid {
		return helpers.ErrUnAuthorizedInValidToken
	}
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return helpers.ErrUnAuthorizedInValidToken
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return helpers.ErrUnAuthorizedTokenExpired
		}
		// if claims["iss"] != "your-issuer" {
		// 	return helpers.ErrUnAuthorizedInvalidIssuer
		// }
	}
	return nil
}

// SetContext ...
func (m *GoMiddleware) SetContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var (
			token    *jwt.Token
			user     *models.User
			student  *models.Student
			clientID database.PID
			ok       bool
			// uniID    = database.Parse(ctx.QueryParam("uni_id"))
			useCache = !cachSkipper(ctx)
			claims   jwt.StandardClaims
			// mainRole *models.Role
		)

		// if !uniID.IsValid() {
		// 	return helpers.Reply(ctx, http.StatusUnauthorized, helpers.ErrorBadRequest, nil, nil)
		// }

		if token, ok = ctx.Get("user").(*jwt.Token); ok {
			helpers.SetToContext(ctx, helpers.HeadersTokenKey, token.Raw)

			if err = helpers.JSONTo(token.Claims, &claims); err != nil {
				return next(ctx)
			}

			if user, err = m.userRepo.GetByID(ctx, database.Parse(claims.Id), useCache); err != nil {
				return helpers.Reply(ctx, http.StatusUnauthorized, helpers.ErrorUnAuthorized, nil, nil)
			}

			if user.Status != models.USER_STATUS_ACTIVE {
				return helpers.Reply(ctx, http.StatusUnauthorized, helpers.ErrUserIsNotActive, nil, nil)
			}

			// mainRole = &user.Role
			// if mainRole == nil || !mainRole.ID.IsValid() {
			// 	return helpers.Reply(ctx, http.StatusUnauthorized, helpers.ErrorInvalidUserRoles, nil, nil)
			// }

			// helpers.SetToContext(ctx, helpers.HeadersMainRole, mainRole)
			// helpers.SetToContext(ctx, helpers.HeadersUserRole, user.Role)
			helpers.SetToContext(ctx, helpers.HeadersUser, user)
			helpers.SetToContext(ctx, helpers.HeadersUserID, user.ID)
		} else if token, ok = ctx.Get("student").(*jwt.Token); ok {
			if student, err = m.studentRepository.GetByID(ctx, database.Parse(claims.Id), useCache); err != nil {
				return helpers.Reply(ctx, http.StatusUnauthorized, helpers.ErrorBadRequest, nil, nil)
			}

			if student.Status != models.StudentStatusActive {
				return helpers.Reply(ctx, http.StatusUnauthorized, helpers.ErrUserIsNotActive, nil, nil)
			}

			// mainRole = student.Roles.GetMainRole()
			// if mainRole == nil || !mainRole.ID.IsValid() {
			// 	return helpers.Reply(ctx, http.StatusUnauthorized, helpers.ErrorInvalidUserRoles, nil, nil)
			// }

			// helpers.SetToContext(ctx, helpers.HeadersMainRole, mainRole)
			// helpers.SetToContext(ctx, helpers.HeadersUserRole, student.Roles)
			helpers.SetToContext(ctx, helpers.HeadersStudent, student)
			helpers.SetToContext(ctx, helpers.HeadersStudentID, student.ID)
		} else {
			return helpers.Reply(ctx, http.StatusUnauthorized, helpers.ErrUnAuthorizedInValidToken, nil, nil)
		}

		helpers.SetToContext(ctx, helpers.HeadersTokenKey, token.Raw)
		if err = helpers.JSONTo(token.Claims, &claims); err != nil {
			return next(ctx)
		}

		helpers.SetHeaderToRequest(ctx, helpers.HeadersClientID, clientID)
		return next(ctx)
	}
}

// CheckAuthorization ...
func (m *GoMiddleware) CheckAuthorization(pathURL string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			if RegisterSkipper(ctx) {
				return next(ctx)
			}
			var (
				headerUser = ctx.Get(helpers.HeadersUser)
				// authorized   bool
				// methodAction string
				// mainRole     *models.Role
			)

			if headerUser == nil {
				return helpers.Reply(ctx, http.StatusForbidden, helpers.ErrorBadRequest, nil, nil)
			}

			user := headerUser.(*models.User)
			if user == nil || !user.ID.IsValid() {
				return helpers.Reply(ctx, http.StatusForbidden, helpers.ErrInvalidUserID, nil, nil)
			}
			// mainRole = &user.
			// if !user.Role.ID.IsValid() {
			// 	return helpers.Reply(ctx, http.StatusForbidden, helpers.ErrorInvalidUserRoles, nil, nil)
			// }

			// if mainRole.Priority < 0 {
			// 	return next(ctx)
			// }

			// uniID := helpers.ContextUniID(ctx)
			// methodAction = ctx.Request().Method

			// request := &models.AuthRules{
			// 	PolymorphicModel: models.PolymorphicModel{
			// 		OwnerType: "roles",
			// 		OwnerID:   mainRole.ID,
			// 	},
			// 	V2: "uni",
			// 	V3: uniID.String(),
			// 	V4: pathURL,
			// 	V5: methodAction,
			// }

			// if authorized = m.authUcecase.AuthEnforce(ctx, *request, false); !authorized {
			// 	return helpers.Reply(ctx, http.StatusForbidden, helpers.ErrorAccessDenied, nil, nil)
			// }

			return next(ctx)
		}
	}
}

func (m *GoMiddleware) ForceUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !helpers.ContextUserID(c).IsValid() {
			return helpers.Reply(c, http.StatusUnauthorized, helpers.ErrInvalidUserID, nil, nil)
		}

		return next(c)
	}
}

func (m *GoMiddleware) ForceStudent(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !helpers.ContextStudentID(c).IsValid() {
			return helpers.Reply(c, http.StatusUnauthorized, helpers.ErrInvalidStudentID, nil, nil)
		}

		return next(c)
	}
}
