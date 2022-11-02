package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whitelist []string = make([]string, 5)

type JWTCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
	IsAdmin bool `json:"is_admin"`
}

type ConfigJWT struct {
	SecretJWT      string
	ExpireDuration int
}

func (cj *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(cj.SecretJWT),
	}
}

// GenerateToken perform generating token and exp from userID
func (cj *ConfigJWT) GenerateToken(userID string, isAdmin bool) string {
	claims := JWTCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cj.ExpireDuration))).Unix(),
		},
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	listedToken, _ := token.SignedString([]byte(cj.SecretJWT))

	whitelist = append(whitelist, listedToken)
	return listedToken
}


// GetUser perform claims all user
func GetUser(c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)

	if isListed := CheckToken(user.Raw); !isListed {
		return nil
	}
	claims := user.Claims.(*JWTCustomClaims)
	return claims
}

// IsWho perform athorized user with id and admin
func IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("id")
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JWTCustomClaims)
		isAdmin := claims.IsAdmin
		id := claims.ID

		if bool(isAdmin) == true {
			return next(c)
		}
		if userID == id{
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}

// IsAdmin perform athorized only admin
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JWTCustomClaims)
		isAdmin := claims.IsAdmin

		if bool(isAdmin) == false {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

// CheckToken perform checking the token in whitelist
func CheckToken(token string) bool {
	for _, listedToken := range whitelist {
		if listedToken == token {
			return true
		}
	}
	return false
}

// logout perform deleting token in whitelist
func Logout(token string) bool {
	for i, listedToken := range whitelist {
		if listedToken == token {
			whitelist = append(whitelist[:i], whitelist[i+1:]...)
		}
	}
	return true
}
