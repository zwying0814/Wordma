package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
	"wordma/config"
	"wordma/server/model"
)

// JwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func LoginGetUserToken(user model.User, key string, ttl int) (string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),                                       // 签发时间
			ExpiresAt: time.Now().Add(time.Second * time.Duration(ttl)).Unix(), // 过期时间
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return t, nil
}

var ErrTokenNotProvided = fmt.Errorf("token not provided")
var ErrTokenUserNotFound = fmt.Errorf("user not found")

func GetTokenByReq(c *fiber.Ctx) string {
	token := c.Query("token")
	if token == "" {
		token = c.FormValue("token")
	}
	if token == "" {
		token = c.Get(fiber.HeaderAuthorization)
		token = strings.TrimPrefix(token, "Bearer ")
	}
	return token
}

func GetJwtDataByReq(c *fiber.Ctx) (JwtCustomClaims, error) {
	token := GetTokenByReq(c)
	if token == "" {
		return JwtCustomClaims{}, ErrTokenNotProvided
	}

	parse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected parse signing method=%v", t.Header["alg"])
		}

		return []byte(config.AppKey), nil // 密钥
	})
	if err != nil {
		return JwtCustomClaims{}, err
	}

	claims := JwtCustomClaims{}
	tmp, _ := json.Marshal(parse.Claims)
	_ = json.Unmarshal(tmp, &claims)

	return claims, nil
}

func GetUserByReq(c *fiber.Ctx) (*model.User, error) {
	var err error
	claims, err := GetJwtDataByReq(c)
	if err != nil {
		return nil, err
	}

	user, err := model.GetUserByID(claims.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrTokenUserNotFound
	}

	return user, nil
}
