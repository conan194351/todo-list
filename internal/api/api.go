package api

import (
	"fmt"
	"github.com/conan194351/todo-list.git/internal/api/routes"
	"github.com/conan194351/todo-list.git/internal/config"
	"github.com/conan194351/todo-list.git/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

type Server struct {
	g   *gin.Engine
	rgw *routes.Impl
	pkg *pkg.Package
}

func NewGinEngine() *gin.Engine {
	return gin.New()
}

func NewApiServer(g *gin.Engine, rgw *routes.Impl, pkg *pkg.Package) *Server {
	return &Server{
		g:   g,
		rgw: rgw,
		pkg: pkg,
	}
}

func (api *Server) Run() {
	gin.SetMode(config.GetConfig().Server.GinMode)
	api.g.Use(gin.Logger())
	api.g.Use(gin.Recovery())
	//api.g.Use(api.pkg.Middlewares.VerifyAccessToken())
	api.rgw.Setup()

	api.g.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	api.g.Use()
	fmt.Println("Server is running on " + config.GetConfig().Server.Server)
	err := api.g.Run(config.GetConfig().Server.Server)
	if err != nil {
		panic(err)
	}
}

var ServerModule = fx.Options(
	pkg.PackageModuleFx,
	routes.GateWayModule,
	fx.Provide(NewApiServer),
	fx.Provide(NewGinEngine),
)
