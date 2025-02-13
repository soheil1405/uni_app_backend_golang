package jwt

import (
	"strconv"
	"time"
	"uni_app/models"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// GenerateToken ...
func GenerateToken(auth map[string]string, user *models.User) (tokenKey string, expTime time.Time, err error) {
	var (
		expire    int64
		jwtSecret = viper.GetString("service.auth.secret")
	)

	if expire, err = strconv.ParseInt(auth["expire"], 10, 64); err != nil {
		return
	}

	expTime = time.Now().Local().Add(time.Hour * time.Duration(expire))

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
