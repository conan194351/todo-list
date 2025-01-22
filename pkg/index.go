package pkg

import (
	"github.com/conan194351/todo-list.git/pkg/database/setup"
	"github.com/conan194351/todo-list.git/pkg/jwt"
	"github.com/conan194351/todo-list.git/pkg/midderwares"
	"github.com/conan194351/todo-list.git/pkg/redis"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Package struct {
	Jwt         jwt.Service
	Middlewares *middlewares.Middleware
	Redis       *redis.Client
	DB          *gorm.DB
}

func NewPackage(
	jwt jwt.Service,
	middlewares *middlewares.Middleware,
	redis *redis.Client,
	db *gorm.DB,
) *Package {
	return &Package{
		Jwt:         jwt,
		Middlewares: middlewares,
		Redis:       redis,
		DB:          db,
	}
}

var PackageModuleFx = fx.Options(
	fx.Provide(jwt.NewJWTService),
	fx.Provide(middlewares.NewMiddleware),
	fx.Provide(redis.NewClient),
	setup.DBConnection,
	fx.Provide(NewPackage),
)
