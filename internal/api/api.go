package api

import (
	"github.com/conan194351/todo-list.git/internal/api/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

type Server struct {
	g   *gin.Engine
	rgw *routes.RoutesImpl
}

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func NewApiServer(g *gin.Engine, rgw *routes.RoutesImpl) *Server {
	return &Server{
		g:   g,
		rgw: rgw,
	}
}

func (api *Server) Run() {
	api.g.Use(gin.Logger())
	api.rgw.Setup()

	api.g.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	err := api.g.Run(":" + "8080")
	if err != nil {
		panic(err)
	}
}

var ServerModule = fx.Options(
	routes.RoutesGateWayModule,
	fx.Provide(NewApiServer),
	fx.Provide(NewGinEngine),
)
