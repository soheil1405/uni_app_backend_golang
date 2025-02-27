package jwt

import (
	"strconv"
	"time"
	"uni_app/models"
	"uni_app/services/env"

	"github.com/golang-jwt/jwt"
)

// GenerateToken ...
func GenerateToken(auth map[string]string, user *models.User) (tokenKey string, expTime time.Time, err error) {
	var (
		expire    int64
		jwtSecret = env.GetString("service.auth.secret")
	)

	if expire, err = strconv.ParseInt(auth["expire"], 10, 64); err != nil {
		return
	}

	expTime = time.Now().Add(time.Second * time.Duration(expire))

	claims := &jwt.StandardClaims{
		Id:        user.ID.String(),
		Subject:   user.UserName,
		ExpiresAt: expTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if tokenKey, err = token.SignedString([]byte(jwtSecret)); err != nil {
		return
	}

	return
}
