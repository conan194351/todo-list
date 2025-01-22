package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/conan194351/todo-list.git/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"strings"

	"github.com/conan194351/todo-list.git/pkg/dto/user"
	"github.com/conan194351/todo-list.git/pkg/errs"
	"github.com/conan194351/todo-list.git/pkg/jwt"
	"github.com/conan194351/todo-list.git/pkg/redis"
)

type Middleware struct {
	jwt   jwt.Service
	redis *redis.Client
}

func NewMiddleware(jwt jwt.Service, redisClient *redis.Client) *Middleware {
	return &Middleware{
		jwt:   jwt,
		redis: redisClient,
	}
}

func (m *Middleware) VerifyAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeaderHeader := c.GetHeader("Authorization")
		if authorizationHeaderHeader == "" {
			response.SetHttpStatusError(c, errs.ErrUnauthorized, "user authorization token not found")
			return
		}
		userID, err := m.checkToken(c, authorizationHeaderHeader)
		if err != nil {
			response.SetHttpStatusError(c, errs.ErrUnauthorized, errs.ErrUnauthorized.String())
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}

func (m *Middleware) VerifyRefreshToken(token string) (string, error) {
	claims, err1 := m.jwt.VerifyJWTToken(token)
	if err1 != nil {
		return "", fmt.Errorf("user authorization token not found")
	}
	userID := claims["data"].(string)
	authUserJson, err := m.redis.Get(userID)
	if err != nil {
		return "", fmt.Errorf("user authorization token not found")
	}
	var authUser user.Tokens
	err = json.Unmarshal([]byte(authUserJson), &authUser)
	if err != nil {
		return "", err
	}
	if token != authUser.RefreshToken {
		return "", fmt.Errorf("user authorization token not found")
	}
	return userID, nil
}

func (m *Middleware) checkToken(c *gin.Context, authorizationHeaderHeader string) (string, error) {
	arr := strings.Split(authorizationHeaderHeader, " ")
	if len(arr) <= 1 {
		return "", fmt.Errorf("user authorization token not found")
	}
	token := arr[1]

	claims, err1 := m.jwt.VerifyJWTToken(token)
	if err1 != nil {
		return "", fmt.Errorf("user authorization token not found")
	}

	userID := claims["data"].(string)
	authToken, err := m.redis.Get(userID)
	var authUser user.Tokens
	err = json.Unmarshal([]byte(authToken), &authUser)
	if err != nil {
		return "", err
	}

	if authToken == "" || token != authUser.AccessToken {
		return "", fmt.Errorf("user authorization token not found")
	}

	return userID, nil
}
