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
	config            *models.Config
}

// InitMiddleware ...
func InitMiddleware(db *gorm.DB, config *models.Config) *GoMiddleware {
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

// RegisterSkipper returns true if path =  /api/v1/register
func RegisterSkipper(ctx echo.Context) bool {
	if strings.Contains(ctx.Path(), "/auth/login") ||
		// strings.Contains(ctx.Path(), "/auth/password") ||
		// // strings.Contains(ctx.Path(), "/auth/signup") ||
		// ctx.Path() == "/api/v1/auth/verifycode" ||
		// ctx.Path() == "/api/v1/auth/register" ||
		// ctx.Path() == "/api/v1/auth/resetpassword" ||
		// ctx.Path() == "/api/v1/healthinfo" ||
		// ctx.Path() == "/api/v1/routes" ||
		// ctx.Path() == "/api/v1/config" ||
		// ctx.Path() == "/api/v1/auth/otp/login" ||
		// ctx.Path() == "/api/v1/auth/otp/sms" ||
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
		return models.ErrUnAuthorizedInValidToken
	}
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return models.ErrUnAuthorizedInValidToken
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return models.ErrUnAuthorizedTokenExpired
		}
		// if claims["iss"] != "your-issuer" {
		// 	return models.ErrUnAuthorizedInvalidIssuer
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
			uniID    = database.Parse(ctx.QueryParam("uni_id"))
			useCache = !cachSkipper(ctx)
			claims   jwt.StandardClaims
			// mainRole *models.Role
		)

		if !uniID.IsValid() {
			return helpers.Reply(ctx, http.StatusUnauthorized, models.ErrorBadRequest, nil, nil)
		}

		if token, ok = ctx.Get("user").(*jwt.Token); !ok {
			if user, err = m.userRepo.GetByID(ctx, database.Parse(claims.Id), useCache); err != nil {
				return helpers.Reply(ctx, http.StatusUnauthorized, models.ErrorBadRequest, nil, nil)
			}

			if user.Status != models.USER_STATUS_ACTIVE {
				return helpers.Reply(ctx, http.StatusUnauthorized, models.ErrUserIsNotActive, nil, nil)
			}

			// mainRole = &user.Role
			// if mainRole == nil || !mainRole.ID.IsValid() {
			// 	return helpers.Reply(ctx, http.StatusUnauthorized, models.ErrorInvalidUserRoles, nil, nil)
			// }

			// helpers.SetToContext(ctx, helpers.HeadersMainRole, mainRole)
			// helpers.SetToContext(ctx, helpers.HeadersUserRole, user.Role)
			helpers.SetToContext(ctx, helpers.HeadersUser, user)
			helpers.SetToContext(ctx, helpers.HeadersUserID, user.ID)
		} else if token, ok = ctx.Get("student").(*jwt.Token); !ok {
			if student, err = m.studentRepository.GetByID(ctx, database.Parse(claims.Id), useCache); err != nil {
				return helpers.Reply(ctx, http.StatusUnauthorized, models.ErrorBadRequest, nil, nil)
			}

			if student.Status != models.USER_STATUS_ACTIVE {
				return helpers.Reply(ctx, http.StatusUnauthorized, models.ErrUserIsNotActive, nil, nil)
			}

			// mainRole = student.Roles.GetMainRole()
			// if mainRole == nil || !mainRole.ID.IsValid() {
			// 	return helpers.Reply(ctx, http.StatusUnauthorized, models.ErrorInvalidUserRoles, nil, nil)
			// }

			// helpers.SetToContext(ctx, helpers.HeadersMainRole, mainRole)
			helpers.SetToContext(ctx, helpers.HeadersUserRole, student.Roles)
			helpers.SetToContext(ctx, helpers.HeadersStudent, student)
			helpers.SetToContext(ctx, helpers.HeadersStudentID, student.ID)
		} else {
			return helpers.Reply(ctx, http.StatusUnauthorized, models.ErrUnAuthorizedInValidToken, nil, nil)
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
				return helpers.Reply(ctx, http.StatusForbidden, models.ErrorBadRequest, nil, nil)
			}

			user := headerUser.(*models.User)
			if user == nil || !user.ID.IsValid() {
				return helpers.Reply(ctx, http.StatusForbidden, models.ErrInvalidUserID, nil, nil)
			}
			// mainRole = &user.
			// if !user.Role.ID.IsValid() {
			// 	return helpers.Reply(ctx, http.StatusForbidden, models.ErrorInvalidUserRoles, nil, nil)
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
			// 	return helpers.Reply(ctx, http.StatusForbidden, models.ErrorAccessDenied, nil, nil)
			// }

			return next(ctx)
		}
	}
}

func (m *GoMiddleware) ForceUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !helpers.ContextUserID(c).IsValid() {
			return helpers.Reply(c, http.StatusUnauthorized, models.ErrInvalidUserID, nil, nil)
		}

		return next(c)
	}
}

func (m *GoMiddleware) ForceStudent(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !helpers.ContextStudentID(c).IsValid() {
			return helpers.Reply(c, http.StatusUnauthorized, models.ErrInvalidStudentID, nil, nil)
		}

		return next(c)
	}
}
